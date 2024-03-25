package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "compress/gzip"
   "encoding/base64"
   "fmt"
   "net/http"
   "time"
)

var Platforms = map[int]string{
   // com.google.android.youtube
   0: "x86",
   // com.sygic.aura
   1: "armeabi-v7a",
   // com.kakaogames.twodin
   2: "arm64-v8a",
}

var Phone = Device{
   Texture: []string{
      // com.instagram.android
      "GL_OES_compressed_ETC1_RGB8_texture",
      // com.kakaogames.twodin
      "GL_KHR_texture_compression_astc_ldr",
   },
   Library: []string{
      "android.test.mock",
      "svc",
      "telephony-ext",
      "android.hidl.manager-V1.0-java",
      "lockagent",
      "voip-common",
      "services",
      "ext",
      "bu",
      "requestsync",
      "telephony-common",
      "com.android.media.remotedisplay",
      "tcmclient",
      "content",
      "locksettings",
      "bmgr",
      "uiautomator",
      "abx",
      "com.android.mediadrm.signer",
      "incident-helper-cmd",
      "uinput",
      "javax.obex",
      "org.lineageos.platform",
      "appwidget",
      "sm",
      "hid",
      "monkey",
      "ims-common",
      "com.android.location.provider",
      "am",
      "android.test.runner",
      "qcom.fmradio",
      "telecom",
      "vr",
      "framework-graphics",
      "org.apache.http.legacy",
      "framework",
      "android.test.base",
      "com.android.nfc_extras",
      "com.android.future.usb.accessory",
      "android.hidl.base-V1.0-java",
   },
   Feature: []string{
      "reqGlEsVersion=0x30002",
      "android.hardware.audio.low_latency",
      "android.hardware.audio.output",
      "android.hardware.audio.pro",
      "android.hardware.biometrics.face",
      "android.hardware.bluetooth",
      "android.hardware.bluetooth_le",
      "android.hardware.camera",
      "android.hardware.camera.any",
      "android.hardware.camera.autofocus",
      "android.hardware.camera.capability.manual_post_processing",
      "android.hardware.camera.capability.manual_sensor",
      "android.hardware.camera.capability.raw",
      "android.hardware.camera.flash",
      "android.hardware.camera.front",
      "android.hardware.camera.level.full",
      "android.hardware.consumerir",
      "android.hardware.faketouch",
      "android.hardware.fingerprint",
      "android.hardware.location",
      "android.hardware.location.gps",
      "android.hardware.location.network",
      "android.hardware.microphone",
      "android.hardware.nfc",
      "android.hardware.nfc.any",
      "android.hardware.nfc.ese",
      "android.hardware.nfc.hce",
      "android.hardware.nfc.hcef",
      "android.hardware.nfc.uicc",
      "android.hardware.opengles.aep",
      "android.hardware.ram.normal",
      "android.hardware.screen.landscape",
      "android.hardware.screen.portrait",
      "android.hardware.security.model.compatible",
      "android.hardware.sensor.accelerometer",
      "android.hardware.sensor.compass",
      "android.hardware.sensor.gyroscope",
      "android.hardware.sensor.light",
      "android.hardware.sensor.proximity",
      "android.hardware.sensor.stepcounter",
      "android.hardware.sensor.stepdetector",
      "android.hardware.telephony",
      "android.hardware.telephony.calling",
      "android.hardware.telephony.cdma",
      "android.hardware.telephony.data",
      "android.hardware.telephony.gsm",
      "android.hardware.telephony.ims",
      "android.hardware.telephony.messaging",
      "android.hardware.telephony.radio.access",
      "android.hardware.telephony.subscription",
      "android.hardware.touchscreen",
      "android.hardware.touchscreen.multitouch",
      "android.hardware.touchscreen.multitouch.distinct",
      "android.hardware.touchscreen.multitouch.jazzhand",
      "android.hardware.usb.accessory",
      "android.hardware.usb.host",
      "android.hardware.vulkan.compute",
      "android.hardware.vulkan.level=1",
      "android.hardware.vulkan.version=4198400",
      "android.hardware.wifi",
      "android.hardware.wifi.aware",
      "android.hardware.wifi.direct",
      "android.hardware.wifi.passpoint",
      "android.hardware.wifi.rtt",
      "android.software.activities_on_secondary_displays",
      "android.software.adoptable_storage",
      "android.software.app_enumeration",
      "android.software.app_widgets",
      "android.software.autofill",
      "android.software.backup",
      "android.software.cant_save_state",
      "android.software.companion_device_setup",
      "android.software.controls",
      "android.software.cts",
      "android.software.device_admin",
      "android.software.file_based_encryption",
      "android.software.freeform_window_management",
      "android.software.game_service",
      "android.software.home_screen",
      "android.software.input_methods",
      "android.software.ipsec_tunnels",
      "android.software.live_wallpaper",
      "android.software.managed_users",
      "android.software.midi",
      "android.software.opengles.deqp.level=132383489",
      "android.software.picture_in_picture",
      "android.software.print",
      "android.software.secure_lock_screen",
      "android.software.securely_removes_users",
      "android.software.sip",
      "android.software.sip.voip",
      "android.software.telecom",
      "android.software.voice_recognizers",
      "android.software.vulkan.deqp.level=132383489",
      "android.software.webview",
      "android.software.window_magnification",
      "android.sofware.nfc.beam",
      "com.google.android.apps.dialer.SUPPORTED",
      "com.google.android.apps.dialer.call_recording_audio",
      "com.google.android.apps.photos.NEXUS_PRELOAD",
      "com.google.android.apps.photos.nexus_preload",
      "com.google.android.feature.ADAPTIVE_CHARGING",
      "com.google.android.feature.AER_OPTIMIZED",
      "com.google.android.feature.D2D_CABLE_MIGRATION_FEATURE",
      "com.google.android.feature.DREAMLINER",
      "com.google.android.feature.EXCHANGE_6_2",
      "com.google.android.feature.GMS_GAME_SERVICE",
      "com.google.android.feature.GOOGLE_BUILD",
      "com.google.android.feature.GOOGLE_EXPERIENCE",
      "com.google.android.feature.GOOGLE_FI_BUNDLED",
      "com.google.android.feature.NEXT_GENERATION_ASSISTANT",
      "com.google.android.feature.PIXEL_2017_EXPERIENCE",
      "com.google.android.feature.PIXEL_2018_EXPERIENCE",
      "com.google.android.feature.PIXEL_EXPERIENCE",
      "com.google.android.feature.QUICK_TAP",
      "com.google.android.feature.TURBO_PRELOAD",
      "com.google.android.feature.WELLBEING",
      "com.google.lens.feature.CAMERA_INTEGRATION",
      "com.google.lens.feature.IMAGE_INTEGRATION",
      "com.nxp.mifare",
      "org.lineageos.android",
      "org.lineageos.globalactions",
      "org.lineageos.hardware",
      "org.lineageos.livedisplay",
      "org.lineageos.profiles",
      "org.lineageos.settings",
      "org.lineageos.trust",
   },
}

type Application struct {
   ID string
   Version uint64
}

type Platform int

func (p Platform) String() string {
   return Platforms[int(p)]
}

type Device struct {
   // developer.android.com/guide/topics/manifest/supports-gl-texture-element
   Texture []string
   // developer.android.com/guide/topics/manifest/uses-library-element
   Library []string
   // developer.android.com/guide/topics/manifest/uses-feature-element
   Feature []string
   // developer.android.com/ndk/guides/abis
   Platform string
}

func compress_gzip(p []byte) ([]byte, error) {
   var b bytes.Buffer
   w := gzip.NewWriter(&b)
   if _, err := w.Write(p); err != nil {
      return nil, err
   }
   if err := w.Close(); err != nil {
      return nil, err
   }
   return b.Bytes(), nil
}

func encode_base64(p []byte) ([]byte, error) {
   var b bytes.Buffer
   w := base64.NewEncoder(base64.URLEncoding, &b)
   if _, err := w.Write(p); err != nil {
      return nil, err
   }
   if err := w.Close(); err != nil {
      return nil, err
   }
   return b.Bytes(), nil
}

func (a Application) APK(config string) string {
   var b []byte
   b = fmt.Append(b, a.ID, "-")
   if config != "" {
      b = fmt.Append(b, config, "-")
   }
   b = fmt.Append(b, a.Version, ".apk")
   return string(b)
}

func (a Application) OBB(role uint64) string {
   var b []byte
   if role >= 1 {
      b = append(b, "patch"...)
   } else {
      b = append(b, "main"...)
   }
   b = fmt.Append(b, ".", a.Version, ".", a.ID, ".obb")
   return string(b)
}

func (p *Platform) Set(s string) error {
   _, err := fmt.Sscan(s, p)
   if err != nil {
      return err
   }
   return nil
}

func authorization(r *http.Request, a AccessToken) {
   r.Header.Set("Authorization", "Bearer " + a.Values.Get("Auth"))
}

func user_agent(r *http.Request, single bool) {
   var b []byte
   // `sdk` is needed for `/fdfe/delivery`
   b = append(b, "Android-Finsky (sdk="...)
   // with `/fdfe/acquire`, requests will be rejected with certain apps, if the
   // device was created with too low a version here:
   b = fmt.Append(b, android_api)
   b = append(b, ",versionCode="...)
   // for multiple APKs just tell the truth. for single APK we have to lie.
   // below value is the last version that works.
   if single {
      b = fmt.Append(b, 80919999)
   } else {
      b = fmt.Append(b, google_play_store)
   }
   b = append(b, ')')
   r.Header.Set("User-Agent", string(b))
}
func x_dfe_device_id(r *http.Request, c Checkin) error {
   id, err := c.DeviceId()
   if err != nil {
      return err
   }
   r.Header.Set("X-DFE-Device-ID", fmt.Sprintf("%x", id))
   return nil
}

// github.com/doug-leith/android-protobuf-decoding/blob/main/decoding_helpers.py
func x_ps_rh(r *http.Request, c Checkin) error {
   id, err := c.DeviceId()
   if err != nil {
      return err
   }
   token, err := func() ([]byte, error) {
      var m protobuf.Message
      m.Add(3, func(m *protobuf.Message) {
         m.AddBytes(1, fmt.Append(nil, id))
         m.Add(2, func(m *protobuf.Message) {
            v := time.Now().UnixMicro()
            m.AddBytes(1, fmt.Append(nil, v))
         })
      })
      b, err := compress_gzip(m.Encode())
      if err != nil {
         return nil, err
      }
      return encode_base64(b)
   }()
   if err != nil {
      return err
   }
   ps_rh, err := func() ([]byte, error) {
      var m protobuf.Message
      m.Add(1, func(m *protobuf.Message) {
         m.AddBytes(1, token)
      })
      b, err := compress_gzip(m.Encode())
      if err != nil {
         return nil, err
      }
      return encode_base64(b)
   }()
   if err != nil {
      return err
   }
   r.Header.Set("X-PS-RH", string(ps_rh))
   return nil
}

const android_api = 31

// developer.android.com/guide/topics/manifest/uses-feature-element#glEsVersion
// the device actually uses 0x30000, but some apps require a higher version:
// com.axis.drawingdesk.v3
// so lets lie for now
const gl_es_version = 0x30001

const google_play_store = 82941300

