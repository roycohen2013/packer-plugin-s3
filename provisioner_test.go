package main

import (
	"testing"
)

func TestSourceFormat(t *testing.T) {
	tests := []struct {
		name      string
		conf      map[string]interface{}
		shouldErr bool
	}{
		{
			name: "valid source single object",
			conf: map[string]interface{}{
				"objects": []Object{
					{
						Source:      "bucket/obj",
						Destination: "/dest",
					},
				},
				"access_key": "key",
				"secret_key": "key",
				"endpoint":   "ams1.vultrobjects.com",
			},
		},
		{
			name: "valid source nested object",
			conf: map[string]interface{}{
				"objects": []Object{
					{
						Source:      "bucket/folder/obj",
						Destination: "/dest",
					},
				},
				"access_key": "key",
				"secret_key": "key",
				"endpoint":   "ams1.vultrobjects.com",
			},
		},
		{
			name: "missing bucket",
			conf: map[string]interface{}{
				"objects": []Object{
					{
						Source:      "/obj",
						Destination: "/dest",
					},
				},
				"access_key": "key",
				"secret_key": "key",
				"endpoint":   "ams1.vultrobjects.com",
			},
			shouldErr: true,
		},
		{
			name: "no slashes",
			conf: map[string]interface{}{
				"objects": []Object{
					{
						Source:      "obj",
						Destination: "/dest",
					},
				},
				"access_key": "key",
				"secret_key": "key",
				"endpoint":   "ams1.vultrobjects.com",
			},
			shouldErr: true,
		},
		{
			name: "missing obj",
			conf: map[string]interface{}{
				"objects": []Object{
					{
						Source:      "obj/",
						Destination: "/dest",
					},
				},
				"access_key": "key",
				"secret_key": "key",
				"endpoint":   "ams1.vultrobjects.com",
			},
			shouldErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				p   = &S3Provisioner{}
				err = p.Prepare(tt.conf)
			)

			if tt.shouldErr && err == nil {
				t.Fatal("expected error but got nil")
				return
			}

			if !tt.shouldErr && err != nil {
				t.Fatalf("expected no error but got %v", err)
			}
		})
	}
}
