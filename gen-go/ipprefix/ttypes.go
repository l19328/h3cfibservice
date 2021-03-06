// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package ipprefix

import (
	"bytes"
	"context"
	"sync"
	"fmt"
	thrift "github.com/facebook/fbthrift-go"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal
var _ = context.Background

var GoUnusedProtection__ int;

type Fbbinary = []byte

func FbbinaryPtr(v Fbbinary) *Fbbinary { return &v }

// Attributes:
//  - Addr
//  - IfName
type BinaryAddress struct {
  Addr Fbbinary `thrift:"addr,1,required" db:"addr" json:"addr"`
  // unused field # 2
  IfName *string `thrift:"ifName,3" db:"ifName" json:"ifName,omitempty"`
}

func NewBinaryAddress() *BinaryAddress {
  return &BinaryAddress{}
}


func (p *BinaryAddress) GetAddr() Fbbinary {
  return p.Addr
}
var BinaryAddress_IfName_DEFAULT string
func (p *BinaryAddress) GetIfName() string {
  if !p.IsSetIfName() {
    return BinaryAddress_IfName_DEFAULT
  }
return *p.IfName
}
func (p *BinaryAddress) IsSetIfName() bool {
  return p.IfName != nil
}

func (p *BinaryAddress) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetAddr bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
      issetAddr = true
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetAddr{
    return thrift.NewProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Addr is not set"));
  }
  return nil
}

func (p *BinaryAddress)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadBinary(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := Fbbinary(v)
  p.Addr = temp
}
  return nil
}

func (p *BinaryAddress)  ReadField3(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.IfName = &v
}
  return nil
}

func (p *BinaryAddress) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("BinaryAddress"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *BinaryAddress) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("addr", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:addr: ", p), err) }
  if err := oprot.WriteBinary(p.Addr); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.addr (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:addr: ", p), err) }
  return err
}

func (p *BinaryAddress) writeField3(oprot thrift.Protocol) (err error) {
  if p.IsSetIfName() {
    if err := oprot.WriteFieldBegin("ifName", thrift.STRING, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:ifName: ", p), err) }
    if err := oprot.WriteString(string(*p.IfName)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.ifName (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:ifName: ", p), err) }
  }
  return err
}

func (p *BinaryAddress) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("BinaryAddress(%+v)", *p)
}

// Attributes:
//  - PrefixAddress
//  - PrefixLength
type IpPrefix struct {
  PrefixAddress *BinaryAddress `thrift:"prefixAddress,1" db:"prefixAddress" json:"prefixAddress"`
  PrefixLength int16 `thrift:"prefixLength,2" db:"prefixLength" json:"prefixLength"`
}

func NewIpPrefix() *IpPrefix {
  return &IpPrefix{
PrefixAddress: NewBinaryAddress(),
}
}

var IpPrefix_PrefixAddress_DEFAULT *BinaryAddress
func (p *IpPrefix) GetPrefixAddress() *BinaryAddress {
  if !p.IsSetPrefixAddress() {
    return IpPrefix_PrefixAddress_DEFAULT
  }
return p.PrefixAddress
}

func (p *IpPrefix) GetPrefixLength() int16 {
  return p.PrefixLength
}
func (p *IpPrefix) IsSetPrefixAddress() bool {
  return p.PrefixAddress != nil
}

func (p *IpPrefix) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *IpPrefix)  ReadField1(iprot thrift.Protocol) error {
  p.PrefixAddress = NewBinaryAddress()
  if err := p.PrefixAddress.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.PrefixAddress), err)
  }
  return nil
}

func (p *IpPrefix)  ReadField2(iprot thrift.Protocol) error {
  if v, err := iprot.ReadI16(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.PrefixLength = v
}
  return nil
}

func (p *IpPrefix) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("IpPrefix"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *IpPrefix) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("prefixAddress", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:prefixAddress: ", p), err) }
  if err := p.PrefixAddress.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.PrefixAddress), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:prefixAddress: ", p), err) }
  return err
}

func (p *IpPrefix) writeField2(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("prefixLength", thrift.I16, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:prefixLength: ", p), err) }
  if err := oprot.WriteI16(int16(p.PrefixLength)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.prefixLength (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:prefixLength: ", p), err) }
  return err
}

func (p *IpPrefix) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("IpPrefix(%+v)", *p)
}

// Attributes:
//  - Dest
//  - Nexthops
type UnicastRoute struct {
  Dest *IpPrefix `thrift:"dest,1,required" db:"dest" json:"dest"`
  Nexthops []*BinaryAddress `thrift:"nexthops,2,required" db:"nexthops" json:"nexthops"`
}

func NewUnicastRoute() *UnicastRoute {
  return &UnicastRoute{
Dest: NewIpPrefix(),
}
}

var UnicastRoute_Dest_DEFAULT *IpPrefix
func (p *UnicastRoute) GetDest() *IpPrefix {
  if !p.IsSetDest() {
    return UnicastRoute_Dest_DEFAULT
  }
return p.Dest
}

func (p *UnicastRoute) GetNexthops() []*BinaryAddress {
  return p.Nexthops
}
func (p *UnicastRoute) IsSetDest() bool {
  return p.Dest != nil
}

func (p *UnicastRoute) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetDest bool = false;
  var issetNexthops bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
      issetDest = true
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
      issetNexthops = true
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetDest{
    return thrift.NewProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Dest is not set"));
  }
  if !issetNexthops{
    return thrift.NewProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Nexthops is not set"));
  }
  return nil
}

func (p *UnicastRoute)  ReadField1(iprot thrift.Protocol) error {
  p.Dest = NewIpPrefix()
  if err := p.Dest.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Dest), err)
  }
  return nil
}

func (p *UnicastRoute)  ReadField2(iprot thrift.Protocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*BinaryAddress, 0, size)
  p.Nexthops =  tSlice
  for i := 0; i < size; i ++ {
    _elem0 := NewBinaryAddress()
    if err := _elem0.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
    }
    p.Nexthops = append(p.Nexthops, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *UnicastRoute) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("UnicastRoute"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *UnicastRoute) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("dest", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:dest: ", p), err) }
  if err := p.Dest.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Dest), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:dest: ", p), err) }
  return err
}

func (p *UnicastRoute) writeField2(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("nexthops", thrift.LIST, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:nexthops: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Nexthops)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.Nexthops {
    if err := v.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
    }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:nexthops: ", p), err) }
  return err
}

func (p *UnicastRoute) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("UnicastRoute(%+v)", *p)
}

