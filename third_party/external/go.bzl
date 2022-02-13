load("@bazel_gazelle//:deps.bzl", "go_repository")

def load_external_go_repositories():
    ########## Server request handling ###############
    go_repository(
        name = "com_github_andybalholm_brotli",
        importpath = "github.com/andybalholm/brotli",
        commit = "1d750214c25205863625bb3eb8190a51b2cef26d",  # Sep 22, 2021
    )

    go_repository(
        name = "com_github_valyala_bytebufferpool",
        importpath = "github.com/valyala/bytebufferpool",
        commit = "18533face0dfe7042f8157bba9010bd7f8df54b1",  # Nov 4, 2020
    )

    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        tag = "v1.14.2",  # Jan 25, 2022
    )

    go_repository(
        name = "com_github_valyala_fasthttp",
        importpath = "github.com/valyala/fasthttp",
        tag = "v1.33.0",
    )

    go_repository(
        name = "com_github_buaazp_fasthttprouter",
        importpath = "github.com/buaazp/fasthttprouter",
        tag = "979d6e516ec324575737805deabe0303794c58bd",  # Jan 9, 2019
    )

    ########## Logging ###############
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        tag = "v1.9.0",  # Jul 15, 2021
    )

    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        tag = "v1.7.0",  # May 6, 2021
    )

    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        tag = "v1.21.0",  # Feb 7, 2022
    )
