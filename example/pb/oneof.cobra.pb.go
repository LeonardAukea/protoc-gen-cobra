// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func OneofClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Oneof"),
		Short: "Oneof service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_OneofFetchCommand(cfg),
		_OneofFetchNestedCommand(cfg),
	)
	return cmd
}

func _OneofFetchCommand(cfg *client.Config) *cobra.Command {
	req := &FetchRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Fetch"),
		Short: "Fetch RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Oneof"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Oneof", "Fetch"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewOneofClient(cc)
				v := &FetchRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Fetch(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	Option1 := &FetchRequest_Option1{}
	cmd.PersistentFlags().StringVar(&Option1.Option1, cfg.FlagNamer("Option1"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option1"), func() { req.Choose = Option1 })
	Option2 := &FetchRequest_Option2{}
	cmd.PersistentFlags().StringVar(&Option2.Option2, cfg.FlagNamer("Option2"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option2"), func() { req.Choose = Option2 })
	Option3 := &FetchRequest_Option3{}
	cmd.PersistentFlags().StringVar(&Option3.Option3, cfg.FlagNamer("Option3"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option3"), func() { req.Choose = Option3 })

	return cmd
}

func _OneofFetchNestedCommand(cfg *client.Config) *cobra.Command {
	req := &FetchNestedRequest{
		L0: &FetchNestedRequest_Outer{
			L1: &FetchNestedRequest_Outer_Middle{
				L2: &FetchNestedRequest_Outer_Middle_Inner{},
			},
		},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("FetchNested"),
		Short: "FetchNested RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Oneof"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Oneof", "FetchNested"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewOneofClient(cc)
				v := &FetchNestedRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.FetchNested(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	L0L1L2Option1 := &FetchNestedRequest_Outer_Middle_Inner_Option1{}
	cmd.PersistentFlags().StringVar(&L0L1L2Option1.Option1, cfg.FlagNamer("L0 L1 L2 Option1"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 L1 L2 Option1"), func() { req.L0.L1.L2.Choose = L0L1L2Option1 })
	L0L1L2Option2 := &FetchNestedRequest_Outer_Middle_Inner_Option2{}
	cmd.PersistentFlags().StringVar(&L0L1L2Option2.Option2, cfg.FlagNamer("L0 L1 L2 Option2"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 L1 L2 Option2"), func() { req.L0.L1.L2.Choose = L0L1L2Option2 })
	L0L1L2Option3 := &FetchNestedRequest_Outer_Middle_Inner_Option3{}
	cmd.PersistentFlags().StringVar(&L0L1L2Option3.Option3, cfg.FlagNamer("L0 L1 L2 Option3"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 L1 L2 Option3"), func() { req.L0.L1.L2.Choose = L0L1L2Option3 })

	return cmd
}
