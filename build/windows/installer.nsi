Outfile "iPingTray-setup.exe"

InstallDir "$PROGRAMFILES\iPingTray"

RequestExecutionLevel admin

Page directory
Page instfiles

Section "Install"
    SetOutPath $INSTDIR

    File /r "ipingtray.exe"

    CreateShortCut "$SMSTARTUP\iPingTray.lnk" "$INSTDIR\ipingtray.exe"
    WriteUninstaller "$INSTDIR\uninstall.exe"
SectionEnd

Section "Uninstall"
    Delete "$INSTDIR\ipingtray.exe"
    Delete "$SMSTARTUP\iPingTray.lnk"
    RMDir /r "$INSTDIR"
SectionEnd
