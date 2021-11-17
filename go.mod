module github.com/juicedata/juicefs

go 1.16

require (
	github.com/Azure/azure-sdk-for-go v11.1.1-beta+incompatible
	github.com/Azure/go-autorest/autorest v0.11.17 // indirect
	github.com/DataDog/zstd v1.4.5
	github.com/armon/go-metrics v0.0.0-20190430140413-ec5e00d3c878 // indirect
	github.com/aws/aws-sdk-go v1.35.20
	github.com/billziss-gh/cgofuse v1.4.0
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dnaeon/go-vcr v1.2.0 // indirect
	github.com/emersion/go-webdav v0.3.0
	github.com/fatih/color v1.10.0 // indirect
	github.com/go-redis/redis/v8 v8.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/gops v0.3.13
	github.com/google/uuid v1.1.2
	github.com/hanwen/go-fuse/v2 v2.1.1-0.20210611132105-24a1dfe6b4f8
	github.com/hashicorp/consul/api v1.11.0
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/go-msgpack v1.1.5 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hungys/go-lz4 v0.0.0-20170805124057-19ff7f07f099
	github.com/juicedata/godaemon v0.0.0-20210629045518-3da5144a127d
	github.com/juju/ratelimit v1.0.1
	github.com/kr/fs v0.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.8.0
	github.com/mattn/go-isatty v0.0.12
	github.com/mattn/go-sqlite3 v2.0.1+incompatible
	github.com/miekg/dns v1.1.35 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pingcap/check v0.0.0-20200212061837-5e12011dc712 // indirect
	github.com/pingcap/errors v0.11.5-0.20201126102027-b0a155152ca3 // indirect
	github.com/pingcap/log v0.0.0-20210317133921-96f4fcab92a4
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.10.0
	github.com/prometheus/client_golang v1.9.0
	github.com/prometheus/client_model v0.2.0
	github.com/satori/uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20190318030020-c3a204f8e965 // indirect
	github.com/urfave/cli/v2 v2.3.0
	github.com/vbauerster/mpb/v7 v7.0.3
	github.com/viki-org/dnscache v0.0.0-20130720023526-c70c1f23c5d8
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/tools v0.1.4 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/tools v0.2.0 // indirect
	xorm.io/xorm v1.0.7
)

replace github.com/minio/minio v0.0.0-20210206053228-97fe57bba92c => github.com/juicedata/minio v0.0.0-20210222051636-e7cabdf948f4

replace github.com/hanwen/go-fuse/v2 v2.1.1-0.20210611132105-24a1dfe6b4f8 => github.com/juicedata/go-fuse/v2 v2.1.1-0.20210926080226-cfe1ec802a7f

replace github.com/golang-jwt/jwt v3.2.2+incompatible => github.com/dgrijalva/jwt-go v3.2.0+incompatible

replace github.com/dgrijalva/jwt-go v3.2.0+incompatible => github.com/golang-jwt/jwt v3.2.2+incompatible
