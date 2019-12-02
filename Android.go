package libRkTeeWeaver

import (
    "android/soong/android"
    "android/soong/cc"
    "fmt"
    "strings"
)

func init() {
    fmt.Println("libRkTeeWeaver want to conditional Compile")
    android.RegisterModuleType("cc_libRkTeeWeaver", DefaultsFactory)
}

func DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, AddOpteeShardLibs)
    return module
}

func AddOpteeShardLibs(ctx android.LoadHookContext) {
    type props struct {
        Shared_libs []string
    }
    var src_fix string = "libRkTeeWeaver."+getOpteeVersion(ctx)
    p := &props{}
    p.Shared_libs = append(p.Shared_libs, src_fix)
    ctx.AppendProperties(p)
}

func getOpteeVersion(ctx android.BaseContext) (string) {
    var optee_version string = "v1"
    if (strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM"),"rk3326")) {
        optee_version = "v2"
    }
    fmt.Println("Optee Version: " + optee_version)
    return optee_version
}
