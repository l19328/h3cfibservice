// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        thrift "github.com/facebook/fbthrift-go"
        "github.com/h3copen/h3cfibservice/gen-go/platform"
)

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "   getAllLinks()")
  fmt.Fprintln(os.Stderr, "   getAllNeighbors()")
  fmt.Fprintln(os.Stderr, "  void addIfaceAddresses(string iface,  addrs)")
  fmt.Fprintln(os.Stderr, "  void removeIfaceAddresses(string iface,  addrs)")
  fmt.Fprintln(os.Stderr, "  void syncIfaceAddresses(string iface, i16 family, i16 scope,  addrs)")
  fmt.Fprintln(os.Stderr, "   getIfaceAddresses(string iface, i16 family, i16 scope)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.Transport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewHTTPPostClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewSocket(thrift.SocketAddr(net.JoinHostPort(host, portStr)))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.ProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := platform.NewSystemServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "getAllLinks":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetAllLinks requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetAllLinks())
    fmt.Print("\n")
    break
  case "getAllNeighbors":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetAllNeighbors requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetAllNeighbors())
    fmt.Print("\n")
    break
  case "addIfaceAddresses":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "AddIfaceAddresses requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg35 := flag.Arg(2)
    mbTrans36 := thrift.NewMemoryBufferLen(len(arg35))
    defer mbTrans36.Close()
    _, err37 := mbTrans36.WriteString(arg35)
    if err37 != nil { 
      Usage()
      return
    }
    factory38 := thrift.NewSimpleJSONProtocolFactory()
    jsProt39 := factory38.GetProtocol(mbTrans36)
    containerStruct1 := platform.NewSystemServiceAddIfaceAddressesArgs()
    err40 := containerStruct1.ReadField2(jsProt39)
    if err40 != nil {
      Usage()
      return
    }
    argvalue1 := containerStruct1.Addrs
    value1 := argvalue1
    fmt.Print(client.AddIfaceAddresses(value0, value1))
    fmt.Print("\n")
    break
  case "removeIfaceAddresses":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "RemoveIfaceAddresses requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg42 := flag.Arg(2)
    mbTrans43 := thrift.NewMemoryBufferLen(len(arg42))
    defer mbTrans43.Close()
    _, err44 := mbTrans43.WriteString(arg42)
    if err44 != nil { 
      Usage()
      return
    }
    factory45 := thrift.NewSimpleJSONProtocolFactory()
    jsProt46 := factory45.GetProtocol(mbTrans43)
    containerStruct1 := platform.NewSystemServiceRemoveIfaceAddressesArgs()
    err47 := containerStruct1.ReadField2(jsProt46)
    if err47 != nil {
      Usage()
      return
    }
    argvalue1 := containerStruct1.Addrs
    value1 := argvalue1
    fmt.Print(client.RemoveIfaceAddresses(value0, value1))
    fmt.Print("\n")
    break
  case "syncIfaceAddresses":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "SyncIfaceAddresses requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err49 := (strconv.Atoi(flag.Arg(2)))
    if err49 != nil {
      Usage()
      return
    }
    argvalue1 := byte(tmp1)
    value1 := argvalue1
    tmp2, err50 := (strconv.Atoi(flag.Arg(3)))
    if err50 != nil {
      Usage()
      return
    }
    argvalue2 := byte(tmp2)
    value2 := argvalue2
    arg51 := flag.Arg(4)
    mbTrans52 := thrift.NewMemoryBufferLen(len(arg51))
    defer mbTrans52.Close()
    _, err53 := mbTrans52.WriteString(arg51)
    if err53 != nil { 
      Usage()
      return
    }
    factory54 := thrift.NewSimpleJSONProtocolFactory()
    jsProt55 := factory54.GetProtocol(mbTrans52)
    containerStruct3 := platform.NewSystemServiceSyncIfaceAddressesArgs()
    err56 := containerStruct3.ReadField4(jsProt55)
    if err56 != nil {
      Usage()
      return
    }
    argvalue3 := containerStruct3.Addrs
    value3 := argvalue3
    fmt.Print(client.SyncIfaceAddresses(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "getIfaceAddresses":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "GetIfaceAddresses requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err58 := (strconv.Atoi(flag.Arg(2)))
    if err58 != nil {
      Usage()
      return
    }
    argvalue1 := byte(tmp1)
    value1 := argvalue1
    tmp2, err59 := (strconv.Atoi(flag.Arg(3)))
    if err59 != nil {
      Usage()
      return
    }
    argvalue2 := byte(tmp2)
    value2 := argvalue2
    fmt.Print(client.GetIfaceAddresses(value0, value1, value2))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}