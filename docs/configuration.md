# 配置

imgproxy适配[Twelve-Factor-App](https://12factor.net/)并且可以使用 `ENV` 环境变量配置。

## URL 加密

imgproxy允许使用key和salt对url进行签名。此功能在默认情况下已禁用，但强烈建议在生产中启用。要启用URL签名检查，请定义key/salt对：

* `IMGPROXY_KEY`: 十六进制编码key;
* `IMGPROXY_SALT`: 十六进制编码 salt;
* `IMGPROXY_SIGNATURE_SIZE`: 编码到Base64之前用于签名的字节数。默认值：32；

通过用逗号（`，`）分隔key和salt，可以指定多个key/salt对。imgproxy将检查每对的URL签名。当您需要在应用程序中以零停机时间更改key/salt对时非常有用。

还可以逐行指定使用十六进制编码的key和salts的文件路径（在开发环境中很有用）：

```bash
imgproxy -keypath /path/to/file/with/key -saltpath /path/to/file/with/salt
```

如果需要快速生成随机key/salt对，可以使用以下代码片段快速生成该key/salt对：

```bash
echo $(xxd -g 2 -l 64 -p /dev/random | tr -d '\n')
```

## 服务

* `IMGPROXY_BIND`: 要监听的地址和端口或Unix套接字。默认值：`:8080`；
* `IMGPROXY_NETWORK`: 要使用的网络。已知的网络是 `tcp`, `tcp4`, `tcp6`, `unix`, 和 `unixpacket`. 默认: `tcp`;
* `IMGPROXY_READ_TIMEOUT`: 读取整个图像请求（包括正文）的最长持续时间（秒）。 默认: `10`;
* `IMGPROXY_WRITE_TIMEOUT`: 写入响应的最长持续时间（秒）。 默认: `10`;
* `IMGPROXY_KEEP_ALIVE_TIMEOUT`: 关闭连接前等待下一个请求的最长持续时间（秒）。当设置为 `0`, keep-alive是disabled. 默认: `10`;
* `IMGPROXY_DOWNLOAD_TIMEOUT`: 下载源映像的最长持续时间（秒）。 默认: `5`;
* `IMGPROXY_CONCURRENCY`: 要同时处理的最大图像请求数。 默认: CPU核数乘以2；
* `IMGPROXY_MAX_CLIENTS`: 同时活动连接的最大数目。 默认: `IMGPROXY_CONCURRENCY * 10`;
* `IMGPROXY_TTL`: 在`Expires`和`Cache Control:max age`HTTP头中发送的持续时间（秒）。 默认: `3600` (1小时);
* `IMGPROXY_CACHE_CONTROL_PASSTHROUGH`: 当`true`和源图像响应包含`Expires`或`Cache Control`标题时，请重用这些标题。 默认: false;
* `IMGPROXY_SO_REUSEPORT`: when `true`, enables `SO_REUSEPORT` socket option (currently on linux and darwin only);
* `IMGPROXY_USER_AGENT`: User-Agent header that will be sent with source image request. 默认: `imgproxy/%current_version`;
* `IMGPROXY_USE_ETAG`: when `true`, enables using [ETag](https://en.wikipedia.org/wiki/HTTP_ETag) HTTP header for HTTP cache control. 默认: false;
* `IMGPROXY_CUSTOM_REQUEST_HEADERS`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> list of custom headers that imgproxy will send while requesting the source image, divided by `\;` (can be redefined by `IMGPROXY_CUSTOM_HEADERS_SEPARATOR`). Example: `X-MyHeader1=Lorem\;X-MyHeader2=Ipsum`;
* `IMGPROXY_CUSTOM_RESPONSE_HEADERS`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> list of custom response headers, divided by `\;` (can be redefined by `IMGPROXY_CUSTOM_HEADERS_SEPARATOR`). Example: `X-MyHeader1=Lorem\;X-MyHeader2=Ipsum`;
* `IMGPROXY_CUSTOM_HEADERS_SEPARATOR`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> string that will be used as a custom headers separator. 默认: `\;`;

## 安全

imgproxy保护你免受所谓的图像炸弹的攻击。以下是您如何指定您认为合理的最大图像分辨率：

* `IMGPROXY_MAX_SRC_RESOLUTION`: 源图像的最大分辨率，以百万像素为单位。实际尺寸较大的图像将被拒绝。 默认: `16.8`;
* `IMGPROXY_MAX_SRC_FILE_SIZE`: 源映像的最大大小（字节）。文件大小较大的图像将被拒绝。 当 `0`时, 文件大小检查被禁用。 默认: `0`;

imgproxy可以处理动画图像（GIF、WebP），但由于此操作非常繁重，默认情况下只处理一个帧。可以使用以下变量增加要处理的动画帧的最大值：

* `IMGPROXY_MAX_ANIMATION_FRAMES`: 要处理的动画图像帧的最大值。 默认: `1`.

**📝注意:** imgproxy在检查源图像分辨率时汇总所有帧分辨率。

imgproxy读取一些字节来检查源图像是否为SVG。默认情况下，它读取的最大值为32KB，但您可以更改：

* `IMGPROXY_MAX_SVG_CHECK_BYTES`: imgproxy将读取以识别SVG的最大字节数。如果imgproxy无法识别SVG，请尝试增加此数字。 默认: `32768` (32KB)

您还可以使用HTTP`authorization`头指定用于在生产环境中启用授权的密钥：

* `IMGPROXY_SECRET`: 授权令牌。如果指定，HTTP请求应包含 `Authorization: Bearer %secret%` 头;

默认情况下，imgproxy不发送CORS头。指定允许的源以启用CORS头：

* `IMGPROXY_ALLOW_ORIGIN`: 设置后，启用具有所提供来源的CORS头。默认情况下禁用CORS标题。

您可以限制允许的源URL:

* `IMGPROXY_ALLOWED_SOURCES`: whitelist of source image URLs prefixes divided by comma. When blank, imgproxy allows all source image URLs. Example: `s3://,https://example.com/,local://`. 默认: blank.

**⚠️Warning:** Be careful when using this config to limit source URL hosts, and always add a trailing slash after the host. Bad: `http://example.com`, good: `http://example.com/`. If you don't add a trailing slash, `http://example.com@baddomain.com` will be an allowed URL but the request will be made to `baddomain.com`.

When you use imgproxy in a development environment, it can be useful to ignore SSL verification:

* `IMGPROXY_IGNORE_SSL_VERIFICATION`: when true, disables SSL verification, so imgproxy can be used in a development environment with self-signed SSL certificates.

Also you may want imgproxy to respond with the same error message that it writes to the log:

* `IMGPROXY_DEVELOPMENT_ERRORS_MODE`: when true, imgproxy will respond with detailed error messages. Not recommended for production because some errors may contain stack trace.

## Compression

* `IMGPROXY_QUALITY`: default quality of the resulting image, percentage. 默认: `80`;
* `IMGPROXY_GZIP_COMPRESSION`: GZip compression level. 默认: `5`.

### Advanced JPEG compression

* `IMGPROXY_JPEG_PROGRESSIVE`: when true, enables progressive JPEG compression. 默认: false;
* `IMGPROXY_JPEG_NO_SUBSAMPLE`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> when true, chrominance subsampling is disabled. This will improve quality at the cost of larger file size. 默认: false;
* `IMGPROXY_JPEG_TRELLIS_QUANT`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> when true, enables trellis quantisation for each 8x8 block. Reduces file size but increases compression time. 默认: false;
* `IMGPROXY_JPEG_OVERSHOOT_DERINGING`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> when true, enables overshooting of samples with extreme values. Overshooting may reduce ringing artifacts from compression, in particular in areas where black text appears on a white background. 默认: false;
* `IMGPROXY_JPEG_OPTIMIZE_SCANS`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> when true, split the spectrum of DCT coefficients into separate scans. Reduces file size but increases compression time. Requires `IMGPROXY_JPEG_PROGRESSIVE` to be true. 默认: false;
* `IMGPROXY_JPEG_QUANT_TABLE`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> quantization table to use. Supported values are:
  * `0`: Table from JPEG Annex K (default);
  * `1`: Flat table;
  * `2`: Table tuned for MSSIM on Kodak image set;
  * `3`: Table from ImageMagick by N. Robidoux;
  * `4`: Table tuned for PSNR-HVS-M on Kodak image set;
  * `5`: Table from Relevance of Human Vision to JPEG-DCT Compression (1992);
  * `6`: Table from DCTune Perceptual Optimization of Compressed Dental X-Rays (1997);
  * `7`: Table from A Visual Detection Model for DCT Coefficient Quantization (1993);
  * `8`: Table from An Improved Detection Model for DCT Coefficient Quantization (1993).

**📝Note:** `IMGPROXY_JPEG_TRELLIS_QUANT`, `IMGPROXY_JPEG_OVERSHOOT_DERINGING`, `IMGPROXY_JPEG_OPTIMIZE_SCANS`, and `IMGPROXY_JPEG_QUANT_TABLE` require libvips to be built with [MozJPEG](https://github.com/mozilla/mozjpeg) since standard libjpeg doesn't support those optimizations.

### Advanced PNG compression

* `IMGPROXY_PNG_INTERLACED`: when true, enables interlaced PNG compression. 默认: false;
* `IMGPROXY_PNG_QUANTIZE`: when true, enables PNG quantization. libvips should be built with libimagequant support. 默认: false;
* `IMGPROXY_PNG_QUANTIZATION_COLORS`: maximum number of quantization palette entries. Should be between 2 and 256. 默认: 256;

### Advanced GIF compression

* `IMGPROXY_GIF_OPTIMIZE_FRAMES`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> when true, enables GIF frames optimization. This may produce a smaller result, but may increase compression time.
* `IMGPROXY_GIF_OPTIMIZE_TRANSPARENCY`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> when true, enables GIF transparency optimization. This may produce a smaller result, but may increase compression time.

## WebP support detection

imgproxy can use the `Accept` HTTP header to detect if the browser supports WebP and use it as the default format. This feature is disabled by default and can be enabled by the following options:

* `IMGPROXY_ENABLE_WEBP_DETECTION`: enables WebP support detection. When the file extension is omitted in the imgproxy URL and browser supports WebP, imgproxy will use it as the resulting format;
* `IMGPROXY_ENFORCE_WEBP`: enables WebP support detection and enforces WebP usage. If the browser supports WebP, it will be used as resulting format even if another extension is specified in the imgproxy URL.

When WebP support detection is enabled, please take care to configure your CDN or caching proxy to take the `Accept` HTTP header into account while caching.

**⚠️Warning:** Headers cannot be signed. This means that an attacker can bypass your CDN cache by changing the `Accept` HTTP headers. Have this in mind when configuring your production caching setup.

## Client Hints support

imgproxy can use the `Width`, `Viewport-Width` or `DPR` HTTP headers to determine default width and DPR options using Client Hints. This feature is disabled by default and can be enabled by the following option:

* `IMGPROXY_ENABLE_CLIENT_HINTS`: enables Client Hints support to determine default width and DPR options. Read [here](https://developers.google.com/web/updates/2015/09/automating-resource-selection-with-client-hints) details about Client Hints.

**⚠️Warning:** Headers cannot be signed. This means that an attacker can bypass your CDN cache by changing the `Width`, `Viewport-Width` or `DPR` HTTP headers. Have this in mind when configuring your production caching setup.

## Watermark

* `IMGPROXY_WATERMARK_DATA`: Base64-encoded image data. You can easily calculate it with `base64 tmp/watermark.png | tr -d '\n'`;
* `IMGPROXY_WATERMARK_PATH`: path to the locally stored image;
* `IMGPROXY_WATERMARK_URL`: watermark image URL;
* `IMGPROXY_WATERMARK_OPACITY`: watermark base opacity;
* `IMGPROXY_WATERMARKS_CACHE_SIZE`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> size of custom watermarks cache. When set to `0`, watermarks cache is disabled. By default 256 watermarks are cached.

Read more about watermarks in the [Watermark](watermark.md) guide.

## Presets

Read about imgproxy presets in the [Presets](presets.md) guide.

There are two ways to define presets:

#### Using an environment variable

* `IMGPROXY_PRESETS`: set of preset definitions, comma-divided. Example: `default=resizing_type:fill/enlarge:1,sharp=sharpen:0.7,blurry=blur:2`. 默认: blank.

#### Using a command line argument

```bash
imgproxy -presets /path/to/file/with/presets
```

The file should contain preset definitions, one per line. Lines starting with `#` are treated as comments. Example:

```
default=resizing_type:fill/enlarge:1

# Sharpen the image to make it look better
sharp=sharpen:0.7

# Blur the image to hide details
blurry=blur:2
```

### Using only presets

imgproxy can be switched into "presets-only mode". In this mode, imgproxy accepts only `preset` option arguments as processing options. Example: `http://imgproxy.example.com/unsafe/thumbnail:blurry:watermarked/plain/http://example.com/images/curiosity.jpg@png`

* `IMGPROXY_ONLY_PRESETS`: disable all URL formats and enable presets-only mode.

## Serving local files

imgproxy can serve your local images, but this feature is disabled by default. To enable it, specify your local filesystem root:

* `IMGPROXY_LOCAL_FILESYSTEM_ROOT`: the root of the local filesystem. Keep empty to disable serving of local files.

Check out the [Serving local files](serving_local_files.md) guide to learn more.

## Serving files from Amazon S3

imgproxy can process files from Amazon S3 buckets, but this feature is disabled by default. To enable it, set `IMGPROXY_USE_S3` to `true`:

* `IMGPROXY_USE_S3`: when `true`, enables image fetching from Amazon S3 buckets. 默认: false;
* `IMGPROXY_S3_ENDPOINT`: custom S3 endpoint to being used by imgproxy.

Check out the [Serving files from S3](serving_files_from_s3.md) guide to learn more.

## Serving files from Google Cloud Storage

imgproxy can process files from Google Cloud Storage buckets, but this feature is disabled by default. To enable it, set `IMGPROXY_GCS_KEY` to the content of Google Cloud JSON key:

* `IMGPROXY_GCS_KEY`: Google Cloud JSON key. When set, enables image fetching from Google Cloud Storage buckets. 默认: blank.

Check out the [Serving files from Google Cloud Storage](serving_files_from_google_cloud_storage.md) guide to learn more.

## New Relic metrics

imgproxy can send its metrics to New Relic. Specify your New Relic license key to activate this feature:

* `IMGPROXY_NEW_RELIC_KEY`: New Relic license key;
* `IMGPROXY_NEW_RELIC_APP_NAME`: New Relic application name. 默认: `imgproxy`.

Check out the [New Relic](new_relic.md) guide to learn more.

## Prometheus metrics

imgproxy can collect its metrics for Prometheus. Specify binding for Prometheus metrics server to activate this feature:

* `IMGPROXY_PROMETHEUS_BIND`: Prometheus metrics server binding. Can't be the same as `IMGPROXY_BIND`. 默认: blank.

Check out the [Prometheus](prometheus.md) guide to learn more.

## Error reporting

imgproxy can report occurred errors to Bugsnag, Honeybadger and Sentry:

* `IMGPROXY_BUGSNAG_KEY`: Bugsnag API key. When provided, enables error reporting to Bugsnag;
* `IMGPROXY_BUGSNAG_STAGE`: Bugsnag stage to report to. 默认: `production`;
* `IMGPROXY_HONEYBADGER_KEY`: Honeybadger API key. When provided, enables error reporting to Honeybadger;
* `IMGPROXY_HONEYBADGER_ENV`: Honeybadger env to report to. 默认: `production`;
* `IMGPROXY_SENTRY_DSN`: Sentry project DSN. When provided, enables error reporting to Sentry;
* `IMGPROXY_SENTRY_ENVIRONMENT`: Sentry environment to report to. 默认: `production`;
* `IMGPROXY_SENTRY_RELEASE`: Sentry release to report to. 默认: `imgproxy/{imgproxy version}`;
* `IMGPROXY_REPORT_DOWNLOADING_ERRORS`: when `true`, imgproxy will report downloading errors. 默认: `true`.

## Log

* `IMGPROXY_LOG_FORMAT`: the log format. The following formats are supported:
  * `pretty`: _(default)_ colored human-readable format;
  * `structured`: machine-readable format;
  * `json`: JSON format;
* `IMGPROXY_LOG_LEVEL`: the log level. The following levels are supported `error`, `warn`, `info` and `debug`. 默认: `info`;

imgproxy can send logs to syslog, but this feature is disabled by default. To enable it, set `IMGPROXY_SYSLOG_ENABLE` to `true`:

* `IMGPROXY_SYSLOG_ENABLE`: when `true`, enables sending logs to syslog;
* `IMGPROXY_SYSLOG_LEVEL`: maximum log level to send to syslog. Known levels are: `crit`, `error`, `warning` and `info`. 默认: `info`;
* `IMGPROXY_SYSLOG_NETWORK`: network that will be used to connect to syslog. When blank, the local syslog server will be used. Known networks are `tcp`, `tcp4`, `tcp6`, `udp`, `udp4`, `udp6`, `ip`, `ip4`, `ip6`, `unix`, `unixgram` and `unixpacket`. 默认: blank;
* `IMGPROXY_SYSLOG_ADDRESS`: address of the syslog service. Not used if `IMGPROXY_SYSLOG_NETWORK` is blank. 默认: blank;
* `IMGPROXY_SYSLOG_TAG`: specific syslog tag. 默认: `imgproxy`;

**📝Note:** imgproxy always uses structured log format for syslog.

## Memory usage tweaks

**⚠️Warning:** It's highly recommended to read [Memory usage tweaks](memory_usage_tweaks.md) guide before changing this settings.

* `IMGPROXY_DOWNLOAD_BUFFER_SIZE`: the initial size (in bytes) of a single download buffer. When zero, initializes empty download buffers. 默认: `0`;
* `IMGPROXY_GZIP_BUFFER_SIZE`: the initial size (in bytes) of a single GZip buffer. When zero, initializes empty GZip buffers. Makes sense only when GZip compression is enabled. 默认: `0`;
* `IMGPROXY_FREE_MEMORY_INTERVAL`: the interval (in seconds) at which unused memory will be returned to the OS. 默认: `10`;
* `IMGPROXY_BUFFER_POOL_CALIBRATION_THRESHOLD`: the number of buffers that should be returned to a pool before calibration. 默认: `1024`.

## 其他

* `IMGPROXY_BASE_URL`: 将添加到每个请求的图像URL的基URL前缀。例如，如果基URL是 `http://example.com/images` 和 `/path/to/image.png` 被请求, imgproxy将从 `http://example.com/images/path/to/image.png`下载源图像。 默认: blank.
* `IMGPROXY_USE_LINEAR_COLORSPACE`: when `true`, imgproxy will process images in linear colorspace. This will slow down processing. Note that images won't be fully processed in linear colorspace while shrink-on-load is enabled (see below).
* `IMGPROXY_DISABLE_SHRINK_ON_LOAD`: when `true`, disables shrink-on-load for JPEG and WebP. Allows to process the whole image in linear colorspace but dramatically slows down resizing and increases memory usage when working with large images.
* `IMGPROXY_APPLY_UNSHARPEN_MASKING`: <img class="pro-badge" src="assets/pro.svg" alt="pro" /> when `true`, imgproxy will apply unsharpen masking to the resulting image if one is smaller than the source. 默认: `true`.
* `IMGPROXY_STRIP_METADATA`: whether to strip all metadata (EXIF, IPTC, etc.) from JPEG and WebP output images. 默认: `true`.
