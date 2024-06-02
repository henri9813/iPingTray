Outfile "iPingTray-setup.exe"

InstallDir "$PROGRAMFILES\iPingTray"

RequestExecutionLevel admin

Page directory
Page instfiles

Section "Install"
    SetOutPath $INSTDIR

    File /r "ipingtray.exe"

    CreateShortCut "$SMSTARTUP\iPingTray.lnk" "$INSTDIR\ipingtray.exe"
    CreateShortCut "$SMPROGRAMS\iPingTray.lnk" "$INSTDIR\ipingtray.exe" "" "$INSTDIR\ipingtray.exe" 0

    WriteUninstaller "$INSTDIR\uninstall.exe"
SectionEnd

Section "Uninstall"
    Delete "$INSTDIR\ipingtray.exe"

    Delete "$SMSTARTUP\iPingTray.lnk"
    Delete "$SMPROGRAMS\iPingTray.lnk"

    RMDir "$INSTDIR"
SectionEnd
