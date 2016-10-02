// Code generated by protoc-gen-cobra.
// source: store.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	log "log"
	grpc "google.golang.org/grpc"
	json "encoding/json"
	template "text/template"
	time "time"
	tls "crypto/tls"
	x509 "crypto/x509"
	iocodec "github.com/fiorix/protoc-gen-cobra/iocodec"
	context "context"
	ioutil "io/ioutil"
	pflag "github.com/spf13/pflag"
	cobra "github.com/spf13/cobra"
	filepath "path/filepath"
	os "os"
	credentials "google.golang.org/grpc/credentials"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ cobra.Command
var _ context.Context
var _ = ioutil.Discard
var _ pflag.FlagSet
var _ credentials.AuthInfo
var _ filepath.WalkFunc
var _ os.File
var _ grpc.ClientConn
var _ log.Logger
var _ tls.Config
var _ x509.Certificate
var _ iocodec.Encoder
var _ json.Encoder
var _ template.Template
var _ time.Time

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

var _DefaultStoreClientCmdConfig = _NewStoreClientCmdConfig()

func init() {
	_DefaultStoreClientCmdConfig.AddFlags(StoreClientCommand.PersistentFlags())
}

type _StoreClientCmdConfig struct {
	ServerAddr         string
	RequestFile        string
	PrintSampleRequest bool
	ResponseFormat     string
	Timeout            time.Duration
	TLS                bool
	InsecureSkipVerify bool
	CACertFile         string
	CertFile           string
	KeyFile            string
}

func _NewStoreClientCmdConfig() *_StoreClientCmdConfig {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = "localhost:8080"
	}
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		timeout = 10 * time.Second
	}
	outfmt := os.Getenv("RESPONSE_FORMAT")
	if outfmt == "" {
		outfmt = "yaml"
	}
	return &_StoreClientCmdConfig{
		ServerAddr:         addr,
		RequestFile:        os.Getenv("REQUEST_FILE"),
		PrintSampleRequest: os.Getenv("PRINT_SAMPLE_REQUEST") != "",
		ResponseFormat:     outfmt,
		Timeout:            timeout,
		TLS:                os.Getenv("TLS") != "",
		InsecureSkipVerify: os.Getenv("TLS_INSECURE_SKIP_VERIFY") != "",
		CACertFile:         os.Getenv("TLS_CA_CERT_FILE"),
		CertFile:           os.Getenv("TLS_CERT_FILE"),
		KeyFile:            os.Getenv("TLS_KEY_FILE"),
	}
}

func (o *_StoreClientCmdConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.ServerAddr, "server-addr", "s", o.ServerAddr, "server address in form of host:port")
	fs.StringVarP(&o.RequestFile, "request-file", "f", o.RequestFile, "client request file (extension must be yaml, json, or xml)")
	fs.BoolVar(&o.PrintSampleRequest, "print-sample-request", o.PrintSampleRequest, "print sample request file and exit")
	fs.StringVarP(&o.ResponseFormat, "response-format", "o", o.ResponseFormat, "response format (yaml, json, or xml)")
	fs.DurationVar(&o.Timeout, "timeout", o.Timeout, "client connection timeout")
	fs.BoolVar(&o.TLS, "tls", o.TLS, "enable tls")
	fs.BoolVar(&o.InsecureSkipVerify, "tls-insecure-skip-verify", o.InsecureSkipVerify, "INSECURE: skip tls checks")
	fs.StringVar(&o.CACertFile, "tls-ca-cert-file", o.CACertFile, "ca certificate file")
	fs.StringVar(&o.CertFile, "tls-cert-file", o.CertFile, "client certificate file")
	fs.StringVar(&o.KeyFile, "tls-key-file", o.KeyFile, "client key file")
}

var StoreClientCommand = &cobra.Command{
	Use: "store",
}

func _DialStore() (*grpc.ClientConn, StoreClient, error) {
	cfg := _DefaultStoreClientCmdConfig
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(cfg.Timeout),
	}
	if cfg.TLS {
		tlsConfig := tls.Config{}
		if cfg.InsecureSkipVerify {
			tlsConfig.InsecureSkipVerify = true
		}
		if cfg.CACertFile != "" {
			cacert, err := ioutil.ReadFile(cfg.CACertFile)
			if err != nil {
				return nil, nil, fmt.Errorf("ca cert: %v", err)
			}
			certpool := x509.NewCertPool()
			certpool.AppendCertsFromPEM(cacert)
			tlsConfig.RootCAs = certpool
		}
		if cfg.CertFile != "" {
			if cfg.KeyFile == "" {
				return nil, nil, fmt.Errorf("missing key file")
			}
			pair, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
			if err != nil {
				return nil, nil, fmt.Errorf("cert/key: %v", err)
			}
			tlsConfig.Certificates = []tls.Certificate{pair}
		}
		cred := credentials.NewTLS(&tlsConfig)
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(cfg.ServerAddr, opts...)
	if err != nil {
		return nil, nil, err
	}
	return conn, NewStoreClient(conn), nil
}

type _StoreRoundTripFunc func(cli StoreClient) (out interface{}, err error)

func _StoreRoundTrip(v interface{}, fn _StoreRoundTripFunc) error {
	cfg := _DefaultStoreClientCmdConfig
	var e iocodec.EncoderMaker
	var ok bool
	if cfg.ResponseFormat == "" {
		e = iocodec.DefaultEncoders["yaml"]
	} else {
		e, ok = iocodec.DefaultEncoders[cfg.ResponseFormat]
		if !ok {
			return fmt.Errorf("invalid response format: %q", cfg.ResponseFormat)
		}
	}
	if cfg.PrintSampleRequest {
		return e.NewEncoder(os.Stdout).Encode(v)
	}
	if cfg.RequestFile == "" {
		return fmt.Errorf("no request file")
	}
	ext := filepath.Ext(cfg.RequestFile)
	if len(ext) > 0 && ext[0] == '.' {
		ext = ext[1:]
	}
	d, ok := iocodec.DefaultDecoders[ext]
	if !ok {
		return fmt.Errorf("invalid request file format: %q", ext)
	}
	f, err := os.Open(cfg.RequestFile)
	if err != nil {
		return fmt.Errorf("request file: %v", err)
	}
	defer f.Close()
	err = d.NewDecoder(f).Decode(v)
	if err != nil {
		return fmt.Errorf("request parser: %v", err)
	}
	conn, client, err := _DialStore()
	if err != nil {
		return err
	}
	defer conn.Close()
	out, err := fn(client)
	if err != nil {
		return err
	}
	return e.NewEncoder(os.Stdout).Encode(out)
}

var _StoreSetClientCmd = &cobra.Command{
	Use: "set",
	Run: func(cmd *cobra.Command, args []string) {
		var in SetRequest
		err := _StoreRoundTrip(&in, func(cli StoreClient) (interface{}, error) {
			return cli.Set(context.Background(), &in)
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	StoreClientCommand.AddCommand(_StoreSetClientCmd)
}

var _StoreGetClientCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		var in GetRequest
		err := _StoreRoundTrip(&in, func(cli StoreClient) (interface{}, error) {
			return cli.Get(context.Background(), &in)
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	StoreClientCommand.AddCommand(_StoreGetClientCmd)
}

var _StoreMultiGetClientCmd = &cobra.Command{
	Use: "multiget",
	Run: func(cmd *cobra.Command, args []string) {
		var in MultiGetRequest
		err := _StoreRoundTrip(&in, func(cli StoreClient) (interface{}, error) {
			return cli.MultiGet(context.Background(), &in)
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	StoreClientCommand.AddCommand(_StoreMultiGetClientCmd)
}
