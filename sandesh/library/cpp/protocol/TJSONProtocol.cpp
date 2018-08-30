/*
 * Copyright (c) 2013 Juniper Networks, Inc. All rights reserved.
 */

#include <cassert>
#include <cctype>
#include <cstdio>
#include <stdexcept>
#include <boost/static_assert.hpp>
#include <boost/algorithm/string/replace.hpp>

#include <base/logging.h>
#include <base/util.h>
#include <base/string_util.h>

#include "TJSONProtocol.h"

using std::string;

#ifdef TJSONPROTOCOL_DEBUG_PRETTY_PRINT
static const string endl = "\n";
#else
static const string endl = "";
#endif // !TJSONPROTOCOL_DEBUG_PRETTY_PRINT

namespace contrail { namespace sandesh { namespace protocol {

// Static data

static const uint8_t kcJSONTagO = '<';
static const uint8_t kcJSONTagC = '>';
static const uint8_t kcJSONSlash = '/';
static const uint8_t kcJSONSBracketC = ']';

static const std::string kJSONTagO("{");
static const std::string kJSONTagC("}");
static const std::string kJSONSoloTagC("/>");
static const std::string kJSONEndTagO("</");
static const std::string kJSONSlash("/");
static const std::string kJSONType("type");
static const std::string kJSONIdentifier("identifier");
static const std::string kJSONName("name");
static const std::string kJSONKey("key");
static const std::string kJSONElementTagO("");
static const std::string kJSONElementTagC("");
static const std::string kJSONValue("value");
static const std::string kJSONSize("size");
static const std::string kJSONBoolTrue("true");
static const std::string kJSONBoolFalse("false");
static const std::string kJSONCDATAO("<![CDATA[");
static const std::string kJSONCDATAC("]]>");

static const std::string kTypeNameBool("bool");
static const std::string kTypeNameByte("byte");
static const std::string kTypeNameI16("i16");
static const std::string kTypeNameI32("i32");
static const std::string kTypeNameI64("i64");
static const std::string kTypeNameU16("u16");
static const std::string kTypeNameU32("u32");
static const std::string kTypeNameU64("u64");
static const std::string kTypeNameIPV4("ipv4");
static const std::string kTypeNameIPADDR("ipaddr");
static const std::string kTypeNameDouble("double");
static const std::string kTypeNameStruct("struct");
static const std::string kTypeNameString("string");
static const std::string kTypeNameXML("xml");
static const std::string kTypeNameUUID("uuid_t");
static const std::string kTypeNameMap("map");
static const std::string kTypeNameList("list");
static const std::string kTypeNameSet("set");
static const std::string kTypeNameSandesh("sandesh");
static const std::string kTypeNameUnknown("unknown");

static const std::string kAttrTypeSandesh("type=\"sandesh\"");
static const std::string kJSONListTagO("<list ");
static const std::string kJSONListTagC("</list>" + endl);
static const std::string kJSONSetTagO("<set ");
static const std::string kJSONSetTagC("</set>" + endl);
static const std::string kJSONMapTagO("<map ");
static const std::string kJSONMapTagC("</map>" + endl);
static const std::string kType("\"TYPE\":");
static const std::string kTYPEStruct("TYPE: STRUCT");
static const std::string kFieldId("\"ID\":");
static const std::string kVal("\"VAL\":");

const std::string& TJSONProtocol::fieldTypeName(TType type) {
  switch (type) {
    case T_BOOL   : return kTypeNameBool   ;
    case T_BYTE   : return kTypeNameByte   ;
    case T_I16    : return kTypeNameI16    ;
    case T_I32    : return kTypeNameI32    ;
    case T_I64    : return kTypeNameI64    ;
    case T_U16    : return kTypeNameU16    ;
    case T_U32    : return kTypeNameU32    ;
    case T_U64    : return kTypeNameU64    ;
    case T_IPV4   : return kTypeNameIPV4   ;
    case T_IPADDR : return kTypeNameIPADDR ;
    case T_DOUBLE : return kTypeNameDouble ;
    case T_STRING : return kTypeNameString ;
    case T_STRUCT : return kTypeNameStruct ;
    case T_MAP    : return kTypeNameMap    ;
    case T_SET    : return kTypeNameSet    ;
    case T_LIST   : return kTypeNameList   ;
    case T_SANDESH: return kTypeNameSandesh;
    case T_XML    : return kTypeNameXML    ;
    case T_UUID   : return kTypeNameUUID   ;
    default: return kTypeNameUnknown;
  }
}

TType TJSONProtocol::getTypeIDForTypeName(const std::string &name) {
  TType result = T_STOP; // Sentinel value
  if (name.length() > 1) {
    switch (name[0]) {
    case 'b':
        switch (name[1]) {
        case 'o':
            result = T_BOOL;
            break;
        case 'y':
            result = T_BYTE;
            break;
        }
        break;
    case 'd':
      result = T_DOUBLE;
      break;
    case 'i':
      switch (name[1]) {
      case '1':
        result = T_I16;
        break;
      case '3':
        result = T_I32;
        break;
      case '6':
        result = T_I64;
        break;
      case 'p':
        switch (name[2]) {
        case 'a':
          result = T_IPADDR;
          break;
        case 'v':
          result = T_IPV4;
          break;
        }
        break;
      }
      break;
    case 'l':
      result = T_LIST;
      break;
    case 'm':
      result = T_MAP;
      break;
    case 's':
        switch (name[1]) {
        case 'a':
            result = T_SANDESH;
            break;
        case 'e':
            result = T_SET;
            break;
        case 't':
            switch (name[3]) {
            case 'i':
                result = T_STRING;
                break;
            case 'u':
                result = T_STRUCT;
                break;
            }
            break;
        }
        break;
    case 'u':
      switch(name[1]) {
        case '1':
          result = T_U16;
          break;
        case '3':
          result = T_U32;
          break;
        case '6':
          result = T_U64;
          break;
        case 'u':
          result = T_UUID;
          break;
      }
      break;
    case 'x':
      result = T_XML;
      break;
    }
  }
  if (result == T_STOP) {
    LOG(ERROR, __func__ << "Unrecognized type: " << name);
  }
  return result;
}

static inline void formJSONAttr(std::string &dest, const std::string& name, 
                               const std::string & value) {
  dest += name;
  dest += "=\"";
  dest += value; 
  dest += "\"";
}

void TJSONProtocol::indentUp() {
#if TJSONPROTOCOL_DEBUG_PRETTY_PRINT
  indent_str_ += string(indent_inc, ' ');
#endif // !TJSONPROTOCOL_DEBUG_PRETTY_PRINT
}

void TJSONProtocol::indentDown() {
#if TJSONPROTOCOL_DEBUG_PRETTY_PRINT
  if (indent_str_.length() < (string::size_type)indent_inc) {
    LOG(ERROR, __func__ << "Indent string length " <<
        indent_str_.length() << " less than indent length " <<
        (string::size_type)indent_inc);
    return;
  }
  indent_str_.erase(indent_str_.length() - indent_inc);
#endif // !TJSONPROTOOL_DEBUG_PRETTY_PRINT
}

// Returns the number of bytes written on success, -1 otherwise
int32_t TJSONProtocol::writePlain(const string& str) {
  if (is_write_enabled_) {
      int ret = trans_->write((uint8_t*)str.data(), str.length());
      if (ret) {
           return -1;
      }
      return str.length();
  } else {
      return 0;
  }
}

// Returns the number of bytes written on success, -1 otherwise
int32_t TJSONProtocol::writeIndented(const string& str) {
  int ret;

  if (is_write_enabled_) {
#if TJSONPROTOCOL_DEBUG_PRETTY_PRINT
      ret = trans_->write((uint8_t*)indent_str_.data(), indent_str_.length());
      if (ret) {
          return -1;
      }
#endif // !TJSONPROTOCOL_DEBUG_PRETTY_PRINT
      ret = trans_->write((uint8_t*)str.data(), str.length());
      if (ret) {
          return -1;
      }
      return indent_str_.length() + str.length();
  } else {
    return 0;
  }
}

int32_t TJSONProtocol::writeMessageBegin(const std::string& name,
                                        const TMessageType messageType,
                                        const int32_t seqid) {
  return 0;
}

int32_t TJSONProtocol::writeMessageEnd() {
  return 0;
}


/*
* struct_name: { \n
*	"type": STRUCT,
*	"id": ID,
*	"val": { \n
*/
int32_t TJSONProtocol::writeStructBegin(const char* name) {
  int32_t size = 0, ret;
  string sname(name);
  string json;
  json.reserve(512);
  collection_name_stack_.push_back(sname);
  /*
  if(is_list_begin_list_.size() > 0 ) {
      if(!(is_list_begin_list_.back())) {
          json += ",";
      } else {
          is_list_begin_list_.pop_back();
          is_list_begin_list_.push_back(false);
      }
  }
  */

  /* If struct appearing in list context */
  if(!is_list_begin_ && in_list_context_) {
     json += ",";
  } else {
     is_list_begin_ = false;
  }

  /*
  if(!sandesh_begin_) {
      json += "{";
  } else {
      sandesh_begin_ = false;
      sandesh_end_ = true;
  }
  */
  json += "{";
  /*
  json += "\"INSTANCE\":";
  json += "\"";
  json += sname;
  json += "\"";
  json += ","; 
  json += endl;
  */
  json += "\"VAL\":";
  json += kJSONTagO; 
  indentUp();
  // Write to transport
  if ((ret = writeIndented(json)) < 0) {
    LOG(ERROR, __func__ << ": " << sname << " FAILED");
    return ret;
  }
  // push a value to is_struct_begin
  is_struct_begin_list_.push_back(true); 
  is_struct_begin_ = true;
  size += ret;
  return size;
}

/*
 * 	}\n
 * }\n
 */
int32_t TJSONProtocol::writeStructEnd() {
  int32_t size = 0, ret; 
  indentDown();
  string json;
  json.reserve(128);
  if(sandesh_end_) {
      json += kJSONTagC;
      json += endl;
      indentDown(); 
  }
  json += kJSONTagC;
  indentDown();
  json += endl; 
  json += kJSONTagC;
  json += endl;
  indentDown();
  // Write to transport
  if ((ret = writeIndented(json)) < 0) {
    LOG(ERROR, __func__ << ": " << json << " FAILED");
    return ret;
  }
  size += ret;
  if(is_struct_begin_list_.size() > 0) {
      is_struct_begin_list_.pop_back();
  }
  collection_name_stack_.pop_back();
  return size;
}

int32_t TJSONProtocol::writeSandeshBegin(const char* name) {
  int32_t size = 0, ret;
  string sname(name);
  string json;
  sandesh_begin_ = true;
  sandesh_end_ = false;
  json.reserve(512);
  json += kJSONTagO;
  json += endl;
  indentUp();
  json += "\"";
  json += name;
  json += "\"";
  json += ":";
  json += kJSONTagO;
  json += endl;
  indentUp();
  is_write_enabled_ = true; 
  // Write to transport
  ret = writeIndented(json);
  if (ret < 0) {
    LOG(ERROR, __func__ << ": " << sname << " FAILED");
    return ret;
  }
  size += ret;
  is_struct_begin_ = true;
  is_write_enabled_ = false;
  indentUp();
  return size;  
}

int32_t TJSONProtocol::writeSandeshEnd() {

  int32_t size = 0, ret;
  indentDown();
  string json;
  json.reserve(128);
  //indentDown();
  json += kJSONTagC;
  json += endl;
  indentDown();
  json += ",";
  json += "\"TIMESTAMP\":";
  std::stringstream ss;
  std::string t_s;
  ss << UTCTimestampUsec();
  ss >> t_s;
  json += t_s;
  json += kJSONTagC;
  json += endl;
  // Write to transport
  is_write_enabled_ = true;
  if ((ret = writeIndented(json)) < 0) {
    LOG(ERROR, __func__ << ": " << json << " FAILED");
    return ret;
  }
  is_write_enabled_ = false;
  size += ret;
  return size;
}

int32_t TJSONProtocol::writeContainerElementBegin() {
  int32_t size = 0, ret;
  indentDown();
  string json;
  json.reserve(128);
  if(!is_list_begin_ && in_list_context_) {
     json += ",";
  } else {
     is_list_begin_ = false;
  }

  if(in_list_context_ && in_primitive_list_context_) {
      if(is_list_elem_string_) {
          json += "\"";
      }
  } 

  if(in_map_context_) {
      if(!in_primitive_list_context_) {
          if(is_map_primitive_) {
              if(!is_map_val_) {
              // If its a map key primitive type
                  if(!is_map_begin_) {
                      json += ",";
                      json += "\"";
                  } else {
                      is_map_begin_ = false;
                      json += "\"";
                  }
              } else {
                     if (is_map_val_string_) {
                         json += "\"";
                     }
              }
          } else {
              // Non primitive map key 
              if(!is_map_begin_) {
                  json += ",";
                  json += "\""; 
              } else {
                  is_map_begin_ = false;
                  json += "\"";
              }
          }
      }
  }

  // toggle
  if(in_map_context_) {
      if(is_map_primitive_) {
          if(!is_map_val_) {
              is_map_val_ = true;
          } else {
              is_map_val_ = false;
          }
      }
  }

  return writeIndented(json);
  //return writeIndented(kJSONElementTagO);
}

int32_t TJSONProtocol::writeContainerElementEnd() {
  int32_t size = 0, ret;
  indentDown();
  string json;
  json.reserve(128);

  if(in_list_context_ && in_primitive_list_context_) {
      if(is_list_elem_string_) {
          json += "\"";
      }
  }
  if(in_map_context_ && !in_primitive_list_context_) {
      if(is_map_primitive_) {
          if(is_map_val_) {
              json += "\":";
          } else {
              if(is_map_val_string_) {
                  json += "\"";
              }
          }
      } else {

          // non primitive map val
          json += "\":";
      } 
  }
  is_data_map_key_ = false;
  //return writeIndented(kJSONElementTagC);
  return writeIndented(json);
}


/*
 *	"field_name" : {\n
 *		"type": {\n
 *		"val": {\n
 *			<indentup>
 */
int32_t TJSONProtocol::writeFieldBegin(const char *name,
                                      const TType fieldType,
                                      const int16_t fieldId,
                                      const std::map<std::string, std::string> *const amap) {
  int32_t size = 0, ret;
  string sname(name);
  string json;
  json.reserve(512);
  if(!is_struct_begin_) {
      json += ",";
  } else {
      is_struct_begin_ = false;
  }

  json += "\"";
  json += sname;
  json += "\"";
  json += ":";
  json += kJSONTagO; // : {
  json += endl;
  indentUp();
  field_type_.push_back(fieldType);
  json += kType; // "TYPE":
  json += "\"";
  json += fieldTypeName(fieldType);
  json += "\"";
  json += ",";
  json += endl;
  /*
  json += kFieldId; // "ID":
  json += integerToString(fieldId);
  json += ",";
  */
  if(sname == "name") {
      name_field_ = true;
      json += roomKey_prefix_; 
  }
  if(amap != NULL) {
      if(amap->find("tags") != amap->end()) {
          json += "\"COLLECTION_NAME\":\"";
          for(int i=0;i< collection_name_stack_.size(); i++) {
              json += collection_name_stack_[i]+":";
          }
          json += sname;
          json += "\",";
      //}
          // append name to roomkey
          json += endl;
          json += "\"ROOM_KEY_PREFIX\":\"";
          json += roomKey_prefix_;
          json += "\",";
          is_write_enabled_ = true;
      } else if( amap->find("key") != amap->end() ) {
          is_write_enabled_ = true;
      }
  }
  
  if(is_write_enabled_) {
      fields_to_be_written_.push_back(name);
  }

  json += "\"VAL\":";
  //If field type is string quote it
  if(fieldType == T_STRING || fieldType == T_IPADDR) {
      is_string_begin_ = true;
      json += "\""; 
  }
  json += endl;
  // Write to transport
  if ((ret = writeIndented(json)) < 0) {
    LOG(ERROR, __func__ << ": " << json << " FAILED");
    return ret;
  }
  size += ret;  
  return size; 
}

int32_t TJSONProtocol::writeFieldEnd() {
  int32_t size = 0, ret;
  string json;
  contrail::sandesh::protocol::TType fieldType = field_type_.back();
  indentDown();
  if(is_string_begin_) {
      is_string_begin_ = false;
      json += "\"";
  }
  json += "}";
  if ((ret = writeIndented(json)) < 0) {
    LOG(ERROR, __func__ << ": " << json << " FAILED");
    return ret;
  }
  field_type_.pop_back();
  if(name_field_) {
     name_field_ = false;
  }
  if(fields_to_be_written_.size() == 1) {
      is_write_enabled_ = false;
  }
  if(fields_to_be_written_.size() !=0 ) {
      fields_to_be_written_.pop_back();
  }
  size += ret; 
  return size;
}

int32_t TJSONProtocol::writeFieldStop() {
  return 0;
}

int32_t TJSONProtocol::writeMapBegin(const TType keyType,
                                    const TType valType,
                                    const uint32_t size) {

  int32_t bsize = 0, ret;
  string json;
  json.reserve(256);
  json += "{";
  json += endl;
  json += "\"KEY\":";
  json += "\"";
  json += fieldTypeName(keyType);
  json += "\"";
  json += ",";
  json += endl;
  json += "\"VALUE\":";
  json += "\"";
  json += fieldTypeName(valType);
  json += "\"";
  json += ",";
  
  json += "\"VAL\":";
  json += "{";

  indentUp();
  // Write to transport
  ret = writeIndented(json);
  if (ret < 0) {
    LOG(ERROR, __func__ << ": Key: " << fieldTypeName(keyType) <<
        " Value: " << fieldTypeName(valType) << " FAILED");
    return ret;
  }

  if(valType == T_MAP || valType == T_STRUCT || valType == T_LIST) {
     in_non_primitive_map_context_ = true;
  } else {
     is_map_primitive_ = true;
     if(valType == T_STRING) {
         is_map_val_string_ = true;
     }
  }


  bsize += ret;
  indentUp();
  is_beginning_of_map = true;
  is_map_begin_ = true;
  in_map_context_ = true;
  return bsize;  

}

int32_t TJSONProtocol::writeMapEnd() {
  int32_t size = 0, ret;
  string json;
  json.reserve(256); 
  indentDown();
  json += "}";
  json += endl;

  indentDown();
  json += "}";
  json += endl;


  in_map_context_ = false;
  is_map_primitive_ = false;
  is_map_val_string_ = false;
  if ((ret = writeIndented(json)) < 0) {
    LOG(ERROR, __func__ << " FAILED");
    return ret;
  }
  size += ret;
  return size;
}

int32_t TJSONProtocol::writeListBegin(const TType elemType,
                                     const uint32_t size) {
  int32_t bsize = 0, ret;
  string json;
  json.reserve(256);
  json += "{";
  json += "\"INSTANCE\":";
  json += "\"";
  json += fieldTypeName(elemType);
  json += "\"";
  json += ",";
  json += "\"SIZE\":";
  json += integerToString(size);
  json += ",";
  json += endl;
  json += "\"VAL\":";
  json += "[";
  json += endl;
  // Write to transport
  ret = writeIndented(json);
  if (ret < 0) {
    LOG(ERROR, __func__ << ": " << fieldTypeName(elemType) <<
        " FAILED");
    return ret;
  }
  is_list_begin_list_.push_back(true);
  if(elemType != T_STRUCT || elemType != T_MAP ) {
      //primitive type
      is_primitive_list_begin_ = true;
      is_first_primitve_list_elem_ = true;
      if(elemType == T_STRING || elemType == T_IPADDR) {
          is_list_elem_string_ = true;
      } 
      in_primitive_list_context_ = true;
      //in_list_context_ = true;
  } else {
      in_non_primitive_list_context_ = true;
  }
  is_list_begin_ = true;
  in_list_context_ = true;
  bsize += ret;
  indentUp();
  return bsize;
}

int32_t TJSONProtocol::writeListEnd() {
  int32_t size = 0, ret;
  string json;
  json.reserve(32);
  indentDown();
  json += "]";
  indentDown();
  json += "}";
  ret = writeIndented(json);
  
  if (ret < 0) {
    LOG(ERROR, __func__ << " FAILED");
    return ret;
  }
  size += ret;
  in_list_context_ = false;
  is_list_begin_list_.pop_back();
  is_first_primitve_list_elem_ = false;
  in_non_primitive_list_context_ = false;
  in_primitive_list_context_ = false;
  is_list_elem_string_ = false;
  return size;
}

int32_t TJSONProtocol::writeSetBegin(const TType elemType,
                                    const uint32_t size) {
  int32_t bsize = 0, ret;
  string xml;
  xml.reserve(256);
  // Form the xml tag
  xml += kJSONSetTagO;
  formJSONAttr(xml, kJSONType, fieldTypeName(elemType));
  xml += " ";
  formJSONAttr(xml, kJSONSize, integerToString(size));
  xml += kJSONTagC;
  xml += endl;
  // Write to transport
  ret = writeIndented(xml);
  if (ret < 0) {
    LOG(ERROR, __func__ << ": " << fieldTypeName(elemType) <<
        " FAILED");
    return ret;
  }
  bsize += ret;
  indentUp();
  return bsize;
}

int32_t TJSONProtocol::writeSetEnd() {
  int32_t size = 0, ret;
  indentDown();
  if ((ret = writeIndented(kJSONSetTagC)) < 0) {
    LOG(ERROR, __func__ << " FAILED");
    return ret;
  }
  size += ret;
  return size;
}

int32_t TJSONProtocol::writeBool(const bool value) {
  int32_t size = 0, ret;
  string json;
  json.reserve(512);
  json += integerToString(value);

  //return writePlain(value ? kJSONBoolTrue : kJSONBoolFalse);
  return writePlain(json);
}

int32_t TJSONProtocol::writeByte(const int8_t byte) {
  return writePlain(integerToString(byte));
}

int32_t TJSONProtocol::writeI16(const int16_t i16) {
  return writePlain(integerToString(i16));
}

int32_t TJSONProtocol::writeI32(const int32_t i32) {
  int32_t size = 0, ret;
  string json;
  json.reserve(512);
  json += integerToString(i32);
  return writePlain(json);
}

int32_t TJSONProtocol::writeI64(const int64_t i64) {
  return writePlain(integerToString(i64));
}

int32_t TJSONProtocol::writeU16(const uint16_t u16) {
  return writePlain(integerToString(u16));
}

int32_t TJSONProtocol::writeU32(const uint32_t u32) {
  int32_t size = 0, ret;
  string json;
  json.reserve(512);
  json += integerToString(u32);
  return writePlain(json);
}

int32_t TJSONProtocol::writeU64(const uint64_t u64) {
  return writePlain(integerToString(u64));
}

int32_t TJSONProtocol::writeIPV4(const uint32_t ip4) {
  return writePlain(integerToString(ip4));
}

int32_t TJSONProtocol::writeIPADDR(const boost::asio::ip::address& ipaddress) {
  return writePlain(ipaddress.to_string());
}

int32_t TJSONProtocol::writeDouble(const double dub) {
  int32_t size = 0, ret;
  string json;
  json.reserve(512);
  json += integerToString(dub);
  return writePlain(json);
}

int32_t TJSONProtocol::writeString(const string& str) {
  int32_t size = 0, ret;
  string json;
  json.reserve(512);
  if(name_field_) {
      roomKey_prefix_ = str;
  }
  /*
  json += "\"VAL\":";
  json += "\"";
  json += str;
  json += "\""; 
  */
  json += str;
  // Escape JSON control characters in the string before writing
  return writePlain(escapeJSONControlChars(json));
}

int32_t TJSONProtocol::writeBinary(const string& str) {
  // XXX Hex?
  return writeString(str);
}

int32_t TJSONProtocol::writeJSON(const string& str) {
  std::string xmlstr;
  xmlstr.reserve(str.length() + kJSONCDATAO.length() + kJSONCDATAC.length());
  xmlstr += kJSONCDATAO;
  xmlstr += str;
  xmlstr += kJSONCDATAC;
  return writePlain(xmlstr);
}

int32_t TJSONProtocol::writeUUID(const boost::uuids::uuid& uuid) {
  const std::string str = boost::uuids::to_string(uuid);
  return writeString(str);
}
/**
 * Reading functions
 */

// Return true if the character ch is in [-+0-9]; false otherwise
static bool isJSONNumeric(uint8_t ch) {
  switch (ch) {
  case '+':
  case '-':
  case '0':
  case '1':
  case '2':
  case '3':
  case '4':
  case '5':
  case '6':
  case '7':
  case '8':
  case '9':
    return true;
  }
  return false;
}

// Read 1 character from the transport trans and verify that it is the
// expected character ch. Returns 1 if does, -1 if not
static int32_t readSyntaxChar(TJSONProtocol::LookaheadReader &reader,
                               uint8_t ch) {
  uint8_t ch2 = reader.read();
  if (ch2 != ch) {
    LOG(ERROR, __func__ << ": Expected \'" << std::string((char *)&ch, 1) <<
        "\'; got \'" << std::string((char *)&ch2, 1) << "\'.");
    return -1;
  }
  return 1;
}

// Reads string from the transport trans and verify that it is the
// expected string str. Returns 1 if does, -1 if not
static int32_t readSyntaxString(TJSONProtocol::LookaheadReader &reader,
                                const std::string str) {
  int32_t result = 0, ret;
  for (std::string::const_iterator it = str.begin(); it != str.end(); it++) {
    if ((ret = readSyntaxChar(reader, *it)) < 0) {
      return ret;
    }
    result += ret;
  }
  return result;
}

// Reads 1 byte and verifies that it matches ch. Returns 1 if it does,
// -1 otherwise
int32_t TJSONProtocol::readJSONSyntaxChar(uint8_t ch) {
  return readSyntaxChar(reader_, ch);
}

// Reads a JSON number or string and interprets it as a double.
int32_t TJSONProtocol::readJSONDouble(double &num) {
  LOG(ERROR, __func__ << ": Not implemented in TJSONProtocol");
  assert(false);
  return -1;
}

// Reads string and verifies that it matches str. Returns 1 if it does,
// -1 otherwise
int32_t TJSONProtocol::readJSONSyntaxString(const std::string &str) {
  return readSyntaxString(reader_, str);
}

int32_t TJSONProtocol::readJSONCDATA(std::string &str) {
  int32_t result = 0, ret;
  str.clear();
  // Read <![CDATA[
  if ((ret = readJSONSyntaxString(kJSONCDATAO)) < 0) {
    return ret;
  }
  result += ret;
  uint8_t ch = 0, ch2 = 0, ch3;
  bool ch_read = false, ch2_read = false, ch3_read = false;
  while (true) {
    if (!ch_read) {
      ch = reader_.peek();
      reader_.read();
      ++result;
      ch_read = true;
    }
    // Did we read ]]> ?
    if (ch == kcJSONSBracketC) {
      if (!ch2_read) {
        ch2 = reader_.peek();
        reader_.read();
        ++result;
        ch2_read = true;
      }
      if (ch2 == kcJSONSBracketC) {
        if (!ch3_read) {
          ch3 = reader_.peek();
          reader_.read();
          ++result;
          ch3_read = true;
        }
        if (ch3 == kcJSONTagC) {
          break;
        } else {
          // Consume ch
          str += ch;
          // Rotate ch2 and ch3 and repeat
          ch = ch2;
          ch2 = ch3;
          ch3_read = false;
          continue;
        }
      } else {
        // Consume ch
        str += ch;
        // Rotate ch2 and repeat
        ch = ch2;
        ch2_read = false;
        continue;
      }
    } else {
      str += ch;
      ch_read = false;
    }
  }
  return result;
}

// Reads a sequence of characters, stopping at the first occurrence of xml
// tag open delimiter which signals end of field and returns the string
// via str.
int32_t TJSONProtocol::readJSONString(std::string &str) {
  int32_t result = 0;
  str.clear();
  while (true) {
    uint8_t ch = reader_.peek();
    if (ch == kcJSONTagO) {
      break;
    }
    reader_.read();
    str += ch;
    ++result;
  }
  return result;
}

// Decodes an JSON tag and returns the string without the xml open
// and close delimiters via str
int32_t TJSONProtocol::readJSONTag(std::string &str, bool endTag) {
  int32_t result = 0, ret;
  uint8_t ch;
  str.clear();
  if ((ret = readJSONSyntaxChar(kcJSONTagO)) < 0) {
    return ret;
  }
  result += ret;
  if (endTag) {
    if ((ret = readJSONSyntaxChar(kcJSONSlash)) < 0) {
      return ret;
    }
    result += ret;
  }
  while (true) {
    ch = reader_.read();
    ++result;
    if (ch == kcJSONTagC) {
      break;
    }
    str += ch;
  }
  return result;
}

// Reads a sequence of characters, stopping at the first one that is not
// a valid numeric character.
int32_t TJSONProtocol::readJSONNumericChars(std::string &str) {
  uint32_t result = 0;
  str.clear();
  while (true) {
    uint8_t ch = reader_.peek();
    if (!isJSONNumeric(ch)) {
      break;
    }
    reader_.read();
    str += ch;
    ++result;
  }
  return result;
}

// Reads a sequence of characters and assembles them into a number,
// returning them via num
template <typename NumberType>
int32_t TJSONProtocol::readJSONInteger(NumberType &num) {
  int32_t result = 0, ret;
  std::string str;
  if ((ret = readJSONNumericChars(str)) < 0) {
    return ret;
  }
  result += ret;
  stringToInteger(str, num);
  return result;
}

int32_t TJSONProtocol::readMessageBegin(std::string& name,
                                       TMessageType& messageType,
                                       int32_t& seqid) {
  LOG(ERROR, __func__ << ": Not implemented in TJSONProtocol");
  assert(false);
  return -1;
}

int32_t TJSONProtocol::readMessageEnd() {
  LOG(ERROR, __func__ << ": Not implemented in TJSONProtocol");
  assert(false);
  return -1;
}

int32_t TJSONProtocol::readSandeshBegin(std::string& name) {
  std::string str;
  int32_t result = 0, ret;
  if ((ret = readJSONTag(str)) < 0) {
    LOG(ERROR, __func__ << ": FAILED");
    return ret;
  }
  result += ret;
  boost::char_separator<char> sep("=\" ");
  tokenizer tokens(str, sep);
  // Extract the field name
  tokenizer::iterator it = tokens.begin();
  name = *it;
  ++it;
  for (; it != tokens.end(); ++it) {
    if (*it == kJSONType) {
      ++it;
      if (kTypeNameSandesh != *it) {
        LOG(ERROR, __func__ << ": Expected " << kTypeNameSandesh <<
            "; got " << *it);
        return -1;
      }
    }
  }
  return result;
}

int32_t TJSONProtocol::readSandeshEnd() {
  std::string name;
  return readJSONTag(name, true);
}

int32_t TJSONProtocol::readStructBegin(std::string& name) {
  return readJSONTag(name);
}

int32_t TJSONProtocol::readStructEnd() {
  std::string name;
  return readJSONTag(name, true);
}

int32_t TJSONProtocol::readContainerElementBegin() {
  std::string name;
  return readJSONTag(name);
}

int32_t TJSONProtocol::readContainerElementEnd() {
  std::string name;
  return readJSONTag(name, true);
}

int32_t TJSONProtocol::readFieldBegin(std::string& name,
                                     TType& fieldType,
                                     int16_t& fieldId) {
  int32_t result = 0, ret;
  // Check if we hit the end of the list
  uint8_t ch = reader_.peek2();
  uint8_t ch1 = reader_.peek2();
  if (ch == kcJSONTagO && ch1 == kcJSONSlash) {
    fieldType = contrail::sandesh::protocol::T_STOP;
    return result;
  }
  std::string str;
  if ((ret = readJSONTag(str)) < 0) {
    LOG(ERROR, __func__ << ": FAILED");
    return ret;
  }
  result += ret;
  boost::char_separator<char> sep("=\" ");
  tokenizer tokens(str, sep);
  // Extract the field name
  tokenizer::iterator it = tokens.begin();
  name = *it;
  ++it;
  for (; it != tokens.end(); ++it) {
    if (*it == kJSONType) {
      ++it;
      fieldType = getTypeIDForTypeName(*it);
    }
    if (*it == kJSONIdentifier) {
      ++it;
      stringToInteger(*it, fieldId);
    }
  }
  return result;
}

int32_t TJSONProtocol::readFieldEnd() {
  string str;
  return readJSONTag(str, true);
}

int32_t TJSONProtocol::readMapBegin(TType& keyType,
                                   TType& valType,
                                   uint32_t& size) {
  int32_t result = 0, ret;
  std::string str;
  if ((ret = readJSONTag(str)) < 0) {
    LOG(ERROR, __func__ << ": FAILED");
    return ret;
  }
  result += ret;
  boost::char_separator<char> sep("=\" ");
  tokenizer tokens(str, sep);
  // Extract the field name
  tokenizer::iterator it = tokens.begin();
  if (*it != kTypeNameMap) {
    LOG(ERROR, __func__ << ": Expected \"" << kTypeNameMap <<
        "\"; got \"" << *it << "\"");
    return -1;
  }
  ++it;
  for (; it != tokens.end(); ++it) {
    if (*it == kJSONKey) {
      ++it;
      keyType = getTypeIDForTypeName(*it);
    }
    if (*it == kJSONValue) {
      ++it;
      valType = getTypeIDForTypeName(*it);
    }
    if (*it == kJSONSize) {
      ++it;
      stringToInteger(*it, size);
    }
  }
  return result;
}

int32_t TJSONProtocol::readMapEnd() {
  std::string str;
  return readJSONTag(str, true);
}

int32_t TJSONProtocol::readListBegin(TType& elemType,
                                    uint32_t& size) {
  int32_t result = 0, ret;
  std::string str;
  if ((ret = readJSONTag(str)) < 0) {
    LOG(ERROR, __func__ << ": FAILED");
    return ret;
  }
  result += ret;
  boost::char_separator<char> sep("=\" ");
  tokenizer tokens(str, sep);
  // Extract the field name
  tokenizer::iterator it = tokens.begin();
  if (*it != kTypeNameList) {
    LOG(ERROR, __func__ << ": Expected \"" << kTypeNameList <<
        "\"; got \"" << *it << "\"");
    return -1;
  }
  ++it;
  for (; it != tokens.end(); ++it) {
    if (*it == kJSONType) {
      ++it;
      elemType = getTypeIDForTypeName(*it);
    }
    if (*it == kJSONSize) {
      ++it;
      stringToInteger(*it, size);
    }
  }
  return result;
}

int32_t TJSONProtocol::readListEnd() {
  std::string str;
  return readJSONTag(str, true);
}

int32_t TJSONProtocol::readSetBegin(TType& elemType,
                                   uint32_t& size) {
  int32_t result = 0, ret;
  std::string str;
  if ((ret = readJSONTag(str)) < 0) {
    LOG(ERROR, __func__ << ": FAILED");
    return ret;
  }
  result += ret;
  boost::char_separator<char> sep("=\" ");
  tokenizer tokens(str, sep);
  // Extract the field name
  tokenizer::iterator it = tokens.begin();
  if (*it != kTypeNameSet) {
    LOG(ERROR, __func__ << ": Expected \"" << kTypeNameSet <<
        "\"; got \"" << *it << "\"");
    return -1;
  }
  ++it;
  for (; it != tokens.end(); ++it) {
    if (*it == kJSONType) {
      ++it;
      elemType = getTypeIDForTypeName(*it);
    }
    if (*it == kJSONSize) {
      ++it;
      stringToInteger(*it, size);
    }
  }
  return result;
}

int32_t TJSONProtocol::readSetEnd() {
  std::string str;
  return readJSONTag(str, true);
}

int32_t TJSONProtocol::readI16(int16_t& i16) {
  return readJSONInteger(i16);
}

int32_t TJSONProtocol::readI32(int32_t& i32) {
  return readJSONInteger(i32);
}

int32_t TJSONProtocol::readI64(int64_t& i64) {
  return readJSONInteger(i64);
}

int32_t TJSONProtocol::readU16(uint16_t& u16) {
  return readJSONInteger(u16);
}

int32_t TJSONProtocol::readU32(uint32_t& u32) {
  return readJSONInteger(u32);
}

int32_t TJSONProtocol::readU64(uint64_t& u64) {
  return readJSONInteger(u64);
}

int32_t TJSONProtocol::readIPV4(uint32_t& ip4) {
  return readJSONInteger(ip4);
}

int32_t TJSONProtocol::readIPADDR(boost::asio::ip::address& ipaddress) {
  int32_t ret;
  std::string str;
  if ((ret = readJSONString(str)) < 0) {
    return ret;
  }
  boost::system::error_code ec;
  ipaddress = boost::asio::ip::address::from_string(str, ec);
  if (ec) {
    return -1;
  }
  return ret;
}

int32_t TJSONProtocol::readDouble(double& dub) {
  return readJSONDouble(dub);
}

int32_t TJSONProtocol::readString(std::string &str) {
  readJSONString(str);
  unescapeJSONControlChars(str);
  return str.size(); 
}

int32_t TJSONProtocol::readBool(bool& value) {
  std::string str;
  int32_t result = 0, ret;
  if ((ret = readJSONString(str)) < 0) {
    return ret;
  }
  result += ret;
  if (str == kJSONBoolTrue) {
    value = true;
  } else if (str == kJSONBoolFalse) {
    value = false;
  } else {
    LOG(ERROR, __func__ << ": Expected \"" << kJSONBoolTrue <<
        "\" or \"" << kJSONBoolFalse << "\"; got \"" << str << "\"");
  }
  return result;
}

int32_t TJSONProtocol::readByte(int8_t& byte) {
  return readJSONInteger(byte);
}

int32_t TJSONProtocol::readBinary(std::string &str) {
  return readJSONString(str);
}

int32_t TJSONProtocol::readJSON(std::string &str) {
  return readJSONCDATA(str);
}

int32_t TJSONProtocol::readUUID(boost::uuids::uuid &uuid) {
  int32_t ret;
  std::string str;
  if ((ret = readJSONString(str)) < 0) {
    return ret;
  }
  std::stringstream ss;
  ss << str;
  ss >> uuid;
  return ret;
}

}}} // contrail::sandesh::protocol
