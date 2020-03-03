# 生成 URL (高级)

本指南介绍了允许使用所有imgproxy功能的高级URL格式。阅读我们的[生成 URL (基础)](generating_the_url_basic_CN.md) 指南，了解与imgproxy 1.x兼容的“基本”URL格式。

## 格式定义

高级URL应该包含签名、处理选项和源URL，如下所示：

```
/%signature/%processing_options/plain/%source_url@%extension
/%signature/%processing_options/%encoded_source_url.%extension
```

查看本指南末尾的[示例](#Example举例)。

### Signature签名

签名保护您的URL不被攻击者修改。强烈建议在生产环境中签署imgproxy url。

设置了[URL签名](configuration.md#url-signature)后，请查看[Signing the URL](signing_the_url.md)指南以了解如何为URL签名。否则，请在此处使用任何字符串。

### Processing options处理选项

处理选项应指定为URL部分除以斜杠（`/`）。处理选项具有以下格式：
```
%option_name:%argument1:%argument2:...:argumentN
```

处理选项列表未定义imgproxy的处理管道。相反，imgproxy已经提供了一个特定的内置图像处理管道，以获得最高性能。请参阅[关于处理管道](about_processing_pipeline.md)指南中的更多信息。

imgproxy支持以下处理选项：

#### Resize调整大小

```
resize:%resizing_type:%width:%height:%enlarge:%extend
rs:%resizing_type:%width:%height:%enlarge:%extend
```

定义[调整大小类型](#resizing-type)、[宽度](#width)、[高度](#height)、[放大](#enlarge)和[延伸](#extend)的元选项。所有参数都是可选的，可以省略以使用其默认值。

#### Size尺寸

```
size:%width:%height:%enlarge:%extend
s:%width:%height:%enlarge:%extend
```

定义[宽度](#width)、[高度](#height)、[放大](#enlarge)和[延伸](#extend)的元选项。所有参数都是可选的，可以省略以使用其默认值。

#### Resizing type调整大小类型

```
resizing_type:%resizing_type
rt:%resizing_type
```

imgproxy支持以下大小调整类型：

* `fit`: 调整图像的大小，同时保持宽高比以适应给定的大小；
* `fill`: 调整图像大小，同时保持宽高比以填充给定大小并裁剪投影部分；
* `auto`: 如果源维度和结果维度具有相同的方向（纵向或横向），imgproxy将使用`fill`。否则，它将使用`fit`。

默认: `fit`

#### 调整大小算法 <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
resizing_algorithm:%algorithm
ra:%algorithm
```

定义imgproxy将用于调整大小的算法。支持的算法有`nearest`, `linear`, `cubic`, `lanczos2`, 和 `lanczos3`(“最近”、“线性”、“立方”、“lanczos2”和“lanczos3”)。

默认: `lanczos3`

#### Width宽

```
width:%width
w:%width
```

定义结果图像的宽度。当设置为`0`时，imgproxy将使用定义的高度和源纵横比计算得到的宽度。
默认: `0`

#### Height高

```
height:%height
h:%height
```

定义结果图像的高度。当设置为`0`时，imgproxy将使用定义的宽度和源纵横比计算得到的高度。

默认: `0`

#### Dpr设备像素比

```
dpr:%dpr
```

设置后，imgproxy将根据HiDPI（Retina）设备的这个因子乘以图像尺寸。该值必须大于0。

默认: `1`

#### Enlarge放大

```
enlarge:%enlarge
el:%enlarge
```

当设置为`1`、`t`或`true`时，如果图像小于给定大小，imgproxy将放大图像。

默认: false

#### Extend扩展

```
extend:%extend:%gravity
ex:%extend:%gravity
```

* 当`extend`设置为`1`、`t`或`true`时，imgproxy将扩展小于给定大小的图像。
* `gravity` _(optional)_ accepts the same values as [gravity](#gravity) option, except `sm`. When `gravity` is not set, imgproxy will use `ce` gravity without offsets.

默认: `false:ce:0:0`

#### Gravity重力

```
gravity:%gravity_type:%x_offset:%y_offset
g:%gravity_type:%x_offset:%y_offset
```

当imgproxy需要切割图像的某些部分时，它由重力引导。

* `no`: 北（上边缘）；
* `so`: 南（底边）；
* `ea`: 东（右边缘）；
* `we`: 西（左边缘）；
* `noea`: 东北（右上角）；
* `nowe`: 西北（左上角）；
* `soea`: 东南（右下角）；
* `sowe`: 西南（左下角）；
* `ce`: 中心；
* `x_offset`, `y_offset` - （选项）指定X轴和Y轴的重力偏移。

默认: `ce:0:0`

**特殊引力**:

* `gravity:sm` - 智能重力。`libvips检测图像中最“有趣”的部分，并将其视为结果图像的中心。补偿在这里不适用；
* `gravity:fp:%x:%y` -焦点重力。`x`和`y`是介于0和1之间的浮点数，用于定义结果图像中心的坐标。将0和1视为`x`的右/左，将`y`的上/下。

#### Crop作用于

```
crop:%width:%height:%gravity
c:%width:%height:%gravity
```

定义要处理的图像区域（在调整大小之前裁剪）。

* `width` 和 `height` 定义区域的大小。当`width`或`height`设置为`0`时，imgproxy将使用源图像的全宽/全高。
* `gravity` _(选项)_ 接受与[重力](#gravity)选项相同的值。未设置“gravity”时，imgproxy将使用[重力](#gravity)选项的值。

#### Trim修剪

```
trim:%threshold
t:%threshold
```

移除周围背景。

* `threshold` - 颜色相似性公差。

**⚠警告:** 修剪需要将图像完全加载到内存中。这将禁用加载时的缩放，并显著增加内存使用量和处理时间。仔细使用它与大图像。

**📝注意:** 不支持修剪动画图像。

#### Quality质量

```
quality:%quality
q:%quality
```

重新定义结果图像的质量百分比。

默认: 环境变量的值。

#### Max Bytes最大字节数

```
max_bytes:%bytes
mb:%bytes
```

设置后，imgproxy会自动降低图像质量，直到图像低于指定的字节数。

**📝注意:** 仅适用于 `jpg`, `webp`, `heic`, 和 `tiff`.

**⚠警告:** 设置`max_bytes`时，imgproxy会多次保存图像以达到指定的图像大小。

默认: 0

#### Background背景

```
background:%R:%G:%B
bg:%R:%G:%B

background:%hex_color
bg:%hex_color
```

设置后，imgproxy将使用指定的颜色填充生成的图像背景。`R`、`G`、和`B`是背景色（0-255）的红色、绿色和蓝色通道值。`hex_color`是颜色的十六进制编码值。在将具有alpha通道的图像转换为JPEG时非常有用。

如果未提供参数，则禁用任何后台操作。

默认: disabled

#### Adjust调整 <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
adjust:%brightness:%contrast:%saturation
a:%brightness:%contrast:%saturation
```

定义[亮度](#brightness)、[对比度](#contrast)和[饱和度](#saturation)的元选项。所有参数都是可选的，可以省略以使用其默认值。

#### Brightness亮度 <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
brightness:%brightness
br:%brightness
```

设置后，imgproxy将调整结果图像的亮度。`brightness是一个介于`-255`和`255`之间的整数。

默认: 0

#### Contrast对比度 <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
contrast:%contrast
co:%contrast
```

当设置时，Imgproxy将调整结果图像的对比度。`contrast`是一个正浮点数，在那里` 1 `保持对比不稳定性。

默认: 1

#### Saturation饱和度 <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
saturation:%saturation
sa:%saturation
```

设置后，imgproxy将调整结果图像的饱和度。`saturation`是一个正浮点数，其中`1`保持饱和度不变。
默认: 1

#### Blur模糊

```
blur:%sigma
bl:%sigma
```

当设置时，IMGPROXY将高斯模糊滤波器应用于结果图像。`sigma`定义imgproxy将使用的掩码的大小。
默认: disabled

#### Sharpen锐化

```
sharpen:%sigma
sh:%sigma
```

设置后，imgproxy将对生成的图像应用锐化过滤器. `sigma` imgproxy将使用的掩码大小。

As an approximate guideline, use 0.5 sigma for 4 pixels/mm (display resolution), 1.0 for 12 pixels/mm and 1.5 for 16 pixels/mm (300 dpi == 12 pixels/mm).    
作为一个近似准则，对于4像素/mm（显示分辨率），使用0.5西格玛；对于12像素/mm，使用1.0西格玛；对于16像素/mm，使用1.5西格玛（300 dpi==12像素/mm）。

默认: disabled

#### Pixelate像素化 <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
pixelate:%size
pix:%size
```

设置后，imgproxy将对结果图像应用像素化过滤器。`size`是像素的大小。

默认: disabled

#### Watermark水印

```
watermark:%opacity:%position:%x_offset:%y_offset:%scale
wm:%opacity:%position:%x_offset:%y_offset:%scale
```

在处理过的图像上添加水印。

* `opacity` - 水印不透明度修改器。最终不透明度计算如下 `base_opacity * opacity`.
* `position` - （选项）指定水印的位置。可用值:
* `ce`: 中心；
* `no`: 北（上边缘）；
* `so`: 南（底边）；
* `ea`: 东（右边缘）；
* `we`: 西（左边缘）；
* `noea`: 东北（右上角）；
* `nowe`: 西北（左上角）；
* `soea`: 东南（右下角）；
* `sowe`: 西南（左下角）；
  * `re`: 复制水印填充整个图像；
* `x_offset`, `y_offset` - (选项) 指定X和Y轴的水印偏移量。不适用于`re`位置；
* `scale` - (选项)定义相对于结果图像大小的水印大小的浮点数。当设置为`0`或忽略时，水印大小不会更改。

默认: disabled

#### Watermark URL水印URL <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
watermark_url:%url
wmu:%url
```

设置后，imgproxy将使用指定URL中的图像作为水印。`url`是自定义水印的Base64编码url。
默认: blank

#### Style风格 <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
style:%style
st:%style
```

设置后，imgproxy将预先发送`<style>`向源svg图像的`<svg>`节点提供内容的节点。`%style`是url安全的Base64编码CSS样式。

默认: blank

#### JPEG 选项 <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
jpeg_options:%progressive:%no_subsample:%trellis_quant:%overshoot_deringing:%optimize_scans:%quant_table
jpgo:%progressive:%no_subsample:%trellis_quant:%overshoot_deringing:%optimize_scans:%quant_table
```

允许重新定义JPEG保存选项。 所有参数的含义与[高级JPEG压缩](configuration.md#advanced-jpeg-compression)配置相同。  所有参数都是可选的，可以省略。

#### PNG 选项 <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
png_options:%png_interlaced:%png_quantize:%png_quantization_colors
pngo:%png_interlaced:%png_quantize:%png_quantization_colors
```

允许重新定义PNG保存选项。 所有参数的含义与[高级PNG压缩](configuration.md#advanced-png-compression) 。 所有参数都是可选的，可以省略。

#### GIF 选项 <img class="pro-badge" src="assets/pro.svg" alt="pro" />

```
gif_options:%gif_optimize_frames:%gif_optimize_transparency
gifo:%gif_optimize_frames:%gif_optimize_transparency
```

允许重新定义GIF保存选项。 所有参数的含义与[高级GIF压缩](configuration.md#advanced-gif-compression) 。 所有参数都是可选的，可以省略。

#### Preset预设

```
preset:%preset_name1:%preset_name2:...:%preset_nameN
pr:%preset_name1:%preset_name2:...:%preset_nameN
```

定义imgproxy要使用的预设列表。您可以在一个URL中根据需要使用任意多个预设。

阅读[预设](presets.md)指南中有关预设的更多信息。

默认: empty

#### Cache buster高速缓冲存储器

```
cachebuster:%string
cb:%string
```

Cache buster不影响图像处理，但它的变化允许绕过CDN、代理服务器和浏览器缓存。当您更改了某些未反映在URL中的内容（如图像质量设置、预设或水印数据）时非常有用。
强烈建议首选`cachebuster`选项而不是URL查询字符串，因为该选项可以正确签名。
默认: empty

#### Filename文件名

```
filename:%string
fn:%string
```

定义`Content Disposition`头的文件名。如果未指定，imgproxy将从源url获取文件名。
默认: empty

#### Format格式

```
format:%extension
f:%extension
ext:%extension
```

指定生成的图像格式。URL 部分的别名 [扩展名](#extension) .

默认: `jpg`

### Source URL源URL

有两种方法可以指定源url：

#### Plain显式

源URL可以按原样提供，由`/plain/`part:

```
/plain/http://example.com/images/curiosity.jpg
```

**📝注意:** 如果sorce URL包含查询字符串或`@`，则需要对其进行转义。

使用显式URL时, 可以在`@`之后指定 [扩展名](#extension) `:

```
/plain/http://example.com/images/curiosity.jpg@png
```

#### Base64 encodedBase64编码

源URL可以用URL安全BASE64编码。编码后的URL可以用`/`分割，以满足您的需要：

```
/aHR0cDovL2V4YW1w/bGUuY29tL2ltYWdl/cy9jdXJpb3NpdHku/anBn
```

使用编码的源URL时，可以在后面指定 [扩展名](#extension)  `.`:

```
/aHR0cDovL2V4YW1w/bGUuY29tL2ltYWdl/cy9jdXJpb3NpdHku/anBn.png
```

### Extension

扩展名指定结果图像的格式。目前，imgproxy只支持`jpg`、`png`、`webp`、`gif`、`ico`、`heic`和`tiff`，它们是最流行和最有用的图像格式。

<img class="pro-badge" src="assets/pro.svg" alt="pro" /> 您还可以使用`mp4`扩展名将动画图像转换为mp4。

**📝注意:** 阅读有关图像格式支持的详细信息 [这里](image_formats_support.md).

可省略延伸部分。在这种情况下，imgproxy将使用源图像格式作为结果格式。如果结果不支持源图像格式，imgproxy将使用“jpg”。您还可以[启用WebP支持检测](configuration.md#webp-support-detection)尽可能将其用作默认结果格式。

## Example举例

签名的imgproxy URL使用`sharp`预设，调整`http://example.com/images/curiosity.jpg`大小以在不放大的情况下使用智能重力填充`300x400`区域，然后将图像转换为`png`：

```
http://imgproxy.example.com/AfrOrF3gWeDA6VOlDG4TzxMv39O7MXnF4CXpKUwGqRM/preset:sharp/resize:fill:300:400:0/gravity:sm/plain/http://example.com/images/curiosity.jpg@png
```

具有快捷方式的同一URL将如下所示：

```
http://imgproxy.example.com/AfrOrF3gWeDA6VOlDG4TzxMv39O7MXnF4CXpKUwGqRM/pr:sharp/rs:fill:300:400:0/g:sm/plain/http://example.com/images/curiosity.jpg@png
```

与Base64编码的源URL相同的URL将如下所示：

```
http://imgproxy.example.com/AfrOrF3gWeDA6VOlDG4TzxMv39O7MXnF4CXpKUwGqRM/pr:sharp/rs:fill:300:400:0/g:sm/aHR0cDovL2V4YW1w/bGUuY29tL2ltYWdl/cy9jdXJpb3NpdHku/anBn.png
```
