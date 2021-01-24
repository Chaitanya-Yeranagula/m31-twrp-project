package liba_defaults

import (
	"android/soong/android"
	"android/soong/cc"
)

func globalFlags(ctx android.BaseContext) []string {
	var cflags []string

	if ctx.AConfig().Getenv("AB_OTA_UPDATER") == "true" {
		cflags = append(cflags, "-DAB_OTA_UPDATER=1")
	}
	return cflags
}

func globalSrcs(ctx android.BaseContext) []string {
	var srcs []string

	if ctx.AConfig().Getenv("TWRP_CUSTOM_KEYBOARD") != "" {
		srcs = append(srcs, ctx.AConfig().Getenv("TWRP_CUSTOM_KEYBOARD"))
	}

	return srcs
}

func globalIncludes(ctx android.BaseContext) []string {
	var includes []string

	if ctx.AConfig().Getenv("TW_INCLUDE_CRYPTO") != "" {
		includes = append(includes, "bootable/recovery/crypto/fscrypt")
	}

	return includes
}

func libAospRecoveryDefaults(ctx android.LoadHookContext) {
	type props struct {
		Target struct {
			Android struct {
				Cflags  []string
				Enabled *bool
			}
		}
		Cflags       []string
		Srcs         []string
		Include_dirs []string
	}

	p := &props{}
	p.Cflags = globalFlags(ctx)
	s := globalSrcs(ctx)
	p.Srcs = s
	i := globalIncludes(ctx)
	p.Include_dirs = i
	ctx.AppendProperties(p)
}

func init() {
	android.RegisterModuleType("libaosprecovery_defaults", libAospRecoveryDefaultsFactory)
}

func libAospRecoveryDefaultsFactory() android.Module {
	module := cc.DefaultsFactory()
	android.AddLoadHook(module, libAospRecoveryDefaults)

	return module
}
