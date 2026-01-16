package client

const ServerLoadAverageThreshold = 30
const RAMConsumptionThreshold = 0.8
const DiskConsumptionThreshold = 0.9
const NetBandwidthThreshold = 0.9
const Host = "srv.msk01.gigacorp.local"
const URL = "http://" + Host + "/stats"
const Interval = 1
const ContentTypeHeader = "text/plain; charset=utf-8"
const ErrorAmountThreshold = 3
const BitsInMBits = 1000000
const BytesInMBytes = 1024 * 1024
