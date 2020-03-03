# 生成 URL (基础)

本指南介绍了简单的URL格式，该格式易于使用，但不支持所有imgproxy功能。此URL格式主要支持与imgproxy 1.x的向后兼容性。请阅读我们的[生成URL（Advanced）](Generating_URL_Advanced.md)指南以了解高级URL格式。

## 格式定义

基本URL应该包含签名、调整参数大小和源URL，如下所示：

```
/%signature/%resizing_type/%width/%height/%gravity/%enlarge/plain/%source_url@%extension
/%signature/%resizing_type/%width/%height/%gravity/%enlarge/%encoded_source_url.%extension
```

查看本指南末尾的[示例](#Example举例)。

### Signature签名

签名保护您的URL不被攻击者修改。强烈建议在生产环境中签署imgproxy url。

设置了[URL签名](configuration.md#url-signature)后，请查看[Signing the URL](signing_the_url.md)指南以了解如何为URL签名。否则，请在此处使用任何字符串。

### Resizing types调整大小类型

imgproxy支持以下大小调整类型：

* `fit`: 调整图像的大小，同时保持宽高比以适应给定的大小；
* `fill`: 调整图像大小，同时保持宽高比以填充给定大小并裁剪投影部分；
* `auto`: 如果源维度和结果维度具有相同的方向（纵向或横向），imgproxy将使用“fill”。否则，它将使用“fit”。

### Width and height宽和高

宽度和高度参数以像素为单位定义结果图像的大小。根据应用的大小调整类型，尺寸可能与请求的尺寸不同。

### Gravity重力

当imgproxy需要切割图像的某些部分时，它由重力引导。支持以下值：

* `no`: 北（上边缘）；
* `so`: 南（底边）；
* `ea`: 东（右边缘）；
* `we`: 西（左边缘）；
* `noea`: 东北（右上角）；
* `nowe`: 西北（左上角）；
* `soea`: 东南（右下角）；
* `sowe`: 西南（左下角）；
* `ce`: 中心；
* `sm`: 聪明。`libvips `检测图像中最“有趣”的部分，并将其视为结果图像的中心；
* `fp:%x:%y` - 焦点。`x和y是介于0和1之间的浮点数，用于描述结果图像中心的坐标。将0和1视为“x”的右/左，将“y”的上/下。

### Enlarge放大

当设置为“1”、“t”或“true”时，如果图像小于给定大小，imgproxy将放大图像。

### Source URLSource URL

有两种方法可以指定源url：

#### Plain显式

源URL可以按原样提供，由`/plain/`part:

```
/plain/http://example.com/images/curiosity.jpg
```

**📝注意:** 如果sorce URL包含查询字符串或“@”，则需要对其进行转义。

When using plain source URL, you can specify the [extension](#extension) after `@`:

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

扩展名指定结果图像的格式。目前，imgproxy只支持“jpg”、“png”、“webp”、“gif”、“ico”、“heic”和“tiff”，它们是最流行和最有用的图像格式。

<img class="pro-badge" src="assets/pro.svg" alt="pro" /> 您还可以使用“mp4”扩展名将动画图像转换为mp4。

**📝注意:** 阅读有关图像格式支持的详细信息 [这里](image_formats_support.md).

可省略延伸部分。在这种情况下，imgproxy将使用源图像格式作为结果格式。如果结果不支持源图像格式，imgproxy将使用“jpg”。您还可以[启用WebP支持检测](configuration.md#webp-support-detection)尽可能将其用作默认结果格式。

## Example举例

签名的imgproxy URL可调整“http://example.com/images/curiority.jpg”大小，以在不放大的情况下使用智能重力填充“300x400”区域，并将图像转换为“png”：

```
http://imgproxy.example.com/AfrOrF3gWeDA6VOlDG4TzxMv39O7MXnF4CXpKUwGqRM/fill/300/400/sm/0/plain/http://example.com/images/curiosity.jpg@png
```

与Base64编码的源URL相同的URL将如下所示：

```
http://imgproxy.example.com/AfrOrF3gWeDA6VOlDG4TzxMv39O7MXnF4CXpKUwGqRM/fill/300/400/sm/0/aHR0cDovL2V4YW1w/bGUuY29tL2ltYWdl/cy9jdXJpb3NpdHku/anBn.png
```
