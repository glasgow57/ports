--- electron/atom/app/atom_main_delegate.cc.orig	2019-12-17 00:40:10 UTC
+++ electron/atom/app/atom_main_delegate.cc
@@ -8,7 +8,7 @@
 #include <memory>
 #include <string>
 
-#if defined(OS_LINUX)
+#if defined(OS_LINUX) || defined(OS_BSD)
 #include <glib.h>  // for g_setenv()
 #endif
 
@@ -204,7 +204,7 @@ bool AtomMainDelegate::BasicStartupComplete(int* exit_
     base::win::PinUser32();
 #endif
 
-#if defined(OS_LINUX)
+#if defined(OS_LINUX) || defined(OS_BSD)
   // Check for --no-sandbox parameter when running as root.
   if (getuid() == 0 && IsSandboxEnabled(command_line))
     LOG(FATAL) << "Running as root without --"
@@ -229,7 +229,7 @@ void AtomMainDelegate::PostEarlyInitialization(bool is
         ui::ResourceBundle::GetSharedInstance().GetLocaleFilePath(locale, true);
     if (!locale_file_path.empty()) {
       custom_locale = locale;
-#if defined(OS_LINUX)
+#if defined(OS_LINUX) || defined(OS_BSD)
       /* When built with USE_GLIB, libcc's GetApplicationLocaleInternal() uses
        * glib's g_get_language_names(), which keys off of getenv("LC_ALL") */
       g_setenv("LC_ALL", custom_locale.c_str(), TRUE);
