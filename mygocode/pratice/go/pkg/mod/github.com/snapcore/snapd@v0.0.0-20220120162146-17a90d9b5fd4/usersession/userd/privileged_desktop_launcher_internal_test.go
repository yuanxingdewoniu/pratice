// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2021 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package userd_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	. "gopkg.in/check.v1"

	"github.com/snapcore/snapd/dirs"
	"github.com/snapcore/snapd/release"
	"github.com/snapcore/snapd/strutil"
	"github.com/snapcore/snapd/testutil"
	"github.com/snapcore/snapd/usersession/userd"
)

type privilegedDesktopLauncherInternalSuite struct {
	testutil.BaseTest
}

var _ = Suite(&privilegedDesktopLauncherInternalSuite{})

var mockFileSystem = []string{
	"/var/lib/snapd/desktop/applications/mir-kiosk-scummvm_mir-kiosk-scummvm.desktop",
	"/var/lib/snapd/desktop/applications/multipass_multipass-gui.desktop",
	"/var/lib/snapd/desktop/applications/cevelop_cevelop.desktop",
	"/var/lib/snapd/desktop/applications/egmde-confined-desktop_egmde-confined-desktop.desktop",
	"/var/lib/snapd/desktop/applications/classic-snap-analyzer_classic-snap-analyzer.desktop",
	"/var/lib/snapd/desktop/applications/vlc_vlc.desktop",
	"/var/lib/snapd/desktop/applications/gnome-calculator_gnome-calculator.desktop",
	"/var/lib/snapd/desktop/applications/mir-kiosk-kodi_mir-kiosk-kodi.desktop",
	"/var/lib/snapd/desktop/applications/gnome-characters_gnome-characters.desktop",
	"/var/lib/snapd/desktop/applications/clion_clion.desktop",
	"/var/lib/snapd/desktop/applications/gnome-system-monitor_gnome-system-monitor.desktop",
	"/var/lib/snapd/desktop/applications/inkscape_inkscape.desktop",
	"/var/lib/snapd/desktop/applications/gnome-logs_gnome-logs.desktop",
	"/var/lib/snapd/desktop/applications/foo-bar/baz.desktop",
	"/var/lib/snapd/desktop/applications/baz/foo-bar.desktop",

	// A desktop file ID provided by a snap may be shadowed by the
	// host system.
	"/usr/share/applications/shadow-test.desktop",
	"/var/lib/snapd/desktop/applications/shadow-test.desktop",
}

var chromiumDesktopFile = `[Desktop Entry]
X-SnapInstanceName=chromium
Version=1.0
Name=Chromium Web Browser
Name[ast]=Restolador web Chromium
Name[bg]=?????? ?????????? Chromium
Name[bn]=??????????????????????????? ???????????? ????????????????????????
Name[bs]=Chromium web preglednik
Name[ca]=Navegador web Chromium
Name[ca@valencia]=Navegador web Chromium
Name[da]=Chromium netbrowser
Name[de]=Chromium-Webbrowser
Name[en_AU]=Chromium Web Browser
Name[eo]=Kromiumo retfoliumilo
Name[es]=Navegador web Chromium
Name[et]=Chromiumi veebibrauser
Name[eu]=Chromium web-nabigatzailea
Name[fi]=Chromium-selain
Name[fr]=Navigateur Web Chromium
Name[gl]=Navegador web Chromium
Name[he]=?????????? ???????????????? ??????????????
Name[hr]=Chromium web preglednik
Name[hu]=Chromium webb??ng??sz??
Name[hy]=Chromium ???????????? ????????????????
Name[ia]=Navigator del web Chromium
Name[id]=Peramban Web Chromium
Name[it]=Browser web Chromium
Name[ja]=Chromium ????????????????????????
Name[ka]=????????? ???????????????????????? Chromium
Name[ko]=Chromium ??? ????????????
Name[kw]=Peurel wias Chromium
Name[ms]=Pelayar Web Chromium
Name[nb]=Chromium nettleser
Name[nl]=Chromium webbrowser
Name[pt_BR]=Navegador de Internet Chromium
Name[ro]=Navigator Internet Chromium
Name[ru]=??????-?????????????? Chromium
Name[sl]=Chromium spletni brskalnik
Name[sv]=Webbl??saren Chromium
Name[ug]=Chromium ????????????????
Name[vi]=Tr??nh duy???t Web Chromium
Name[zh_CN]=Chromium ???????????????
Name[zh_HK]=Chromium ???????????????
Name[zh_TW]=Chromium ???????????????
GenericName=Web Browser
GenericName[ar]=?????????? ????????????
GenericName[ast]=Restolador web
GenericName[bg]=?????? ??????????????
GenericName[bn]=???????????? ????????????????????????
GenericName[bs]=Web preglednik
GenericName[ca]=Navegador web
GenericName[ca@valencia]=Navegador web
GenericName[cs]=WWW prohl????e??
GenericName[da]=Browser
GenericName[de]=Web-Browser
GenericName[el]=???????????????????? ??????????
GenericName[en_AU]=Web Browser
GenericName[en_GB]=Web Browser
GenericName[eo]=Retfoliumilo
GenericName[es]=Navegador web
GenericName[et]=Veebibrauser
GenericName[eu]=Web-nabigatzailea
GenericName[fi]=WWW-selain
GenericName[fil]=Web Browser
GenericName[fr]=Navigateur Web
GenericName[gl]=Navegador web
GenericName[gu]=????????? ?????????????????????
GenericName[he]=?????????? ??????????????
GenericName[hi]=????????? ????????????????????????
GenericName[hr]=Web preglednik
GenericName[hu]=Webb??ng??sz??
GenericName[hy]=???????????? ????????????????
GenericName[ia]=Navigator del Web
GenericName[id]=Peramban Web
GenericName[it]=Browser web
GenericName[ja]=????????????????????????
GenericName[ka]=????????? ????????????????????????
GenericName[kn]=????????? ??????????????????
GenericName[ko]=??? ????????????
GenericName[kw]=Peurel wias
GenericName[lt]=??iniatinklio nar??ykl??
GenericName[lv]=T??mek??a p??rl??ks
GenericName[ml]=???????????? ????????????????????????
GenericName[mr]=????????? ?????????????????????
GenericName[ms]=Pelayar Web
GenericName[nb]=Nettleser
GenericName[nl]=Webbrowser
GenericName[or]=??????????????? ?????????????????????
GenericName[pl]=Przegl??darka WWW
GenericName[pt]=Navegador Web
GenericName[pt_BR]=Navegador web
GenericName[ro]=Navigator de Internet
GenericName[ru]=??????-??????????????
GenericName[sk]=WWW prehliada??
GenericName[sl]=Spletni brskalnik
GenericName[sr]=???????????????? ????????????????????
GenericName[sv]=Webbl??sare
GenericName[ta]=???????????? ???????????????
GenericName[te]=??????????????? ?????????????????????
GenericName[th]=?????????????????????????????????????????????
GenericName[tr]=Web Taray??c??
GenericName[ug]=????????????????
GenericName[uk]=?????????????????? ??????????
GenericName[vi]=B??? duy???t Web
GenericName[zh_CN]=???????????????
GenericName[zh_HK]=???????????????
GenericName[zh_TW]=???????????????
Comment=Access the Internet
Comment[ar]=???????????? ?????? ????????????????
Comment[ast]=Accesu a Internet
Comment[bg]=???????????? ???? ????????????????
Comment[bn]=?????????????????????????????? ?????????????????? ????????????
Comment[bs]=Pristup internetu
Comment[ca]=Accediu a Internet
Comment[ca@valencia]=Accediu a Internet
Comment[cs]=P????stup k internetu
Comment[da]=F?? adgang til internettet
Comment[de]=Internetzugriff
Comment[el]=???????????????? ?????? ??????????????????
Comment[en_AU]=Access the Internet
Comment[en_GB]=Access the Internet
Comment[eo]=Akiri interreton
Comment[es]=Acceda a Internet
Comment[et]=P????s Internetti
Comment[eu]=Sartu Internetera
Comment[fi]=K??yt?? interneti??
Comment[fil]=I-access ang Internet
Comment[fr]=Acc??der ?? Internet
Comment[gl]=Acceda a Internet
Comment[gu]=????????????????????? ?????????????????? ?????????
Comment[he]=???????? ????????????????
Comment[hi]=????????????????????? ?????? ??????????????? ????????????????????? ????????????
Comment[hr]=Pristupite Internetu
Comment[hu]=Az internet el??r??se
Comment[hy]=?????????? ????????????????
Comment[ia]=Accede a le Interrete
Comment[id]=Akses Internet
Comment[it]=Accesso a Internet
Comment[ja]=????????????????????????????????????
Comment[ka]=?????????????????????????????? ??????????????????
Comment[kn]=??????????????????????????? ??????????????? ???????????????????????????
Comment[ko]=???????????? ???????????????
Comment[kw]=Hedhes an Kesrosweyth
Comment[lt]=Interneto prieiga
Comment[lv]=Piek????t internetam
Comment[ml]=?????????????????????????????????????????? ????????????????????? ?????????????????????
Comment[mr]=???????????????????????????????????? ?????????????????? ?????????
Comment[ms]=Mengakses Internet
Comment[nb]=Bruk internett
Comment[nl]=Verbinding maken met internet
Comment[or]=?????????????????????????????? ?????????????????? ??????????????????
Comment[pl]=Skorzystaj z internetu
Comment[pt]=Aceder ?? Internet
Comment[pt_BR]=Acessar a internet
Comment[ro]=Accesa??i Internetul
Comment[ru]=???????????? ?? ????????????????
Comment[sk]=Pr??stup do siete Internet
Comment[sl]=Dostop do interneta
Comment[sr]=???????????????????? ??????????????????
Comment[sv]=Surfa p?? Internet
Comment[ta]=???????????????????????? ????????????????????????
Comment[te]=???????????????????????????????????? ????????????????????? ????????????????????????
Comment[th]=?????????????????????????????????????????????????????????
Comment[tr]=??nternet'e eri??in
Comment[ug]=?????????????????? ????????????????
Comment[uk]=???????????? ???? ??????????????????
Comment[vi]=Truy c???p Internet
Comment[zh_CN]=???????????????
Comment[zh_HK]=?????????????????????
Comment[zh_TW]=?????????????????????
Exec=env BAMF_DESKTOP_FILE_HINT=/var/lib/snapd/desktop/applications/chromium_chromium.desktop /snap/bin/chromium %U
Terminal=false
Type=Application
Icon=/snap/chromium/1193/chromium.png
Categories=Network;WebBrowser;
MimeType=text/html;text/xml;application/xhtml_xml;x-scheme-handler/http;x-scheme-handler/https;
StartupNotify=true
StartupWMClass=chromium
Actions=NewWindow;Incognito;TempProfile;

[Desktop Action NewWindow]
Name=Open a New Window
Name[ast]=Abrir una Ventana Nueva
Name[bg]=???????????????? ???? ?????? ????????????????
Name[bn]=???????????? ???????????? ?????????????????? ???????????????
Name[bs]=Otvori novi prozor
Name[ca]=Obre una finestra nova
Name[ca@valencia]=Obri una finestra nova
Name[da]=??bn et nyt vindue
Name[de]=Ein neues Fenster ??ffnen
Name[en_AU]=Open a New Window
Name[eo]=Malfermi novan fenestron
Name[es]=Abrir una ventana nueva
Name[et]=Ava uus aken
Name[eu]=Ireki leiho berria
Name[fi]=Avaa uusi ikkuna
Name[fr]=Ouvrir une nouvelle fen??tre
Name[gl]=Abrir unha nova xanela
Name[he]=?????????? ???????? ??????
Name[hy]=?????????? ?????? ????????????????
Name[ia]=Aperi un nove fenestra
Name[it]=Apri una nuova finestra
Name[ja]=?????????????????????????????????
Name[ka]=??????????????? ????????????????????? ??????????????????
Name[kw]=Egery fenester noweth
Name[ms]=Buka Tetingkap Baru
Name[nb]=??pne et nytt vindu
Name[nl]=Nieuw venster openen
Name[pt_BR]=Abre uma nova janela
Name[ro]=Deschide o fereastr?? nou??
Name[ru]=?????????????? ?????????? ????????
Name[sl]=Odpri novo okno
Name[sv]=??ppna ett nytt f??nster
Name[ug]=???????? ???????????? ??????
Name[uk]=???????????????? ???????? ??????????
Name[vi]=M??? c???a s??? m???i
Name[zh_CN]=???????????????
Name[zh_TW]=???????????????
Exec=env BAMF_DESKTOP_FILE_HINT=/var/lib/snapd/desktop/applications/chromium_chromium.desktop /snap/bin/chromium

[Desktop Action Incognito]
Name=Open a New Window in incognito mode
Name[ast]=Abrir una ventana nueva en mou inc??gnitu
Name[bg]=???????????????? ???? ?????? ???????????????? ?? ?????????? \"??????????????????\"
Name[bn]=???????????? ???????????? ?????????????????? ??????????????? ??????????????????????????? ?????????????????????
Name[bs]=Otvori novi prozor u privatnom modu
Name[ca]=Obre una finestra nova en mode d'inc??gnit
Name[ca@valencia]=Obri una finestra nova en mode d'inc??gnit
Name[de]=Ein neues Fenster im Inkognito-Modus ??ffnen
Name[en_AU]=Open a New Window in incognito mode
Name[eo]=Malfermi novan fenestron nekoni??eble
Name[es]=Abrir una ventana nueva en modo inc??gnito
Name[et]=Ava uus aken tundmatus olekus
Name[eu]=Ireki leiho berria isileko moduan
Name[fi]=Avaa uusi ikkuna incognito-tilassa
Name[fr]=Ouvrir une nouvelle fen??tre en mode navigation priv??e
Name[gl]=Abrir unha nova xanela en modo de inc??gnito
Name[he]=?????????? ???????? ?????? ???????? ?????????? ????????
Name[hy]=?????????? ?????? ???????????????? ???????????? ??????????????????????????
Name[ia]=Aperi un nove fenestra in modo incognite
Name[it]=Apri una nuova finestra in modalit?? incognito
Name[ja]=??????????????????????????? ????????????????????????
Name[ka]=??????????????? ????????????????????? ?????????????????????????????? ??????????????????
Name[kw]=Egry fenester noweth en modh privedh
Name[ms]=Buka Tetingkap Baru dalam mod menyamar
Name[nl]=Nieuw venster openen in incognito-modus
Name[pt_BR]=Abrir uma nova janela em modo an??nimo
Name[ro]=Deschide o fereastr?? nou?? ??n mod incognito
Name[ru]=?????????????? ?????????? ???????? ?? ???????????? ??????????????????
Name[sl]=Odpri novo okno v na??inu brez bele??enja
Name[sv]=??ppna ett nytt inkognitof??nster
Name[ug]=?????????????? ?????????????? ???????? ???????????? ??????
Name[uk]=???????????????? ???????? ?????????? ?? ???????????????????? ????????????
Name[vi]=M??? c???a s??? m???i trong ch??? ????? ???n danh
Name[zh_CN]=??????????????????????????????
Name[zh_TW]=??????????????????????????????
Exec=env BAMF_DESKTOP_FILE_HINT=/var/lib/snapd/desktop/applications/chromium_chromium.desktop /snap/bin/chromium --incognito

[Desktop Action TempProfile]
Name=Open a New Window with a temporary profile
Name[ast]=Abrir una ventana nueva con perfil temporal
Name[bg]=???????????????? ???? ?????? ???????????????? ?? ???????????????? ????????????
Name[bn]=?????????????????? ???????????????????????? ?????? ???????????? ???????????? ?????????????????? ???????????????
Name[bs]=Otvori novi prozor pomo??u privremenog profila
Name[ca]=Obre una finestra nova amb un perfil temporal
Name[ca@valencia]=Obri una finestra nova amb un perfil temporal
Name[de]=Ein neues Fenster mit einem tempor??ren Profil ??ffnen
Name[en_AU]=Open a New Window with a temporary profile
Name[eo]=Malfermi novan fenestron portempe
Name[es]=Abrir una ventana nueva con perfil temporal
Name[et]=Ava uus aken ajutise profiiliga
Name[eu]=Ireki leiho berria behin-behineko profil batekin
Name[fi]=Avaa uusi ikkuna k??ytt??en v??liaikaista profiilia
Name[fr]=Ouvrir une nouvelle fen??tre avec un profil temporaire
Name[gl]=Abrir unha nova xanela con perfil temporal
Name[he]=?????????? ???????? ?????? ???? ???????????? ????????
Name[hy]=?????????? ?????? ???????????????? ?????????????????????? ??????????????????
Name[ia]=Aperi un nove fenestra con un profilo provisori
Name[it]=Apri una nuova finestra con un profilo temporaneo
Name[ja]=????????????????????????????????????????????????????????????
Name[ka]=??????????????? ????????????????????? ?????????????????? ????????????????????? ????????????????????????
Name[kw]=Egery fenester noweth gen profil dres prys
Name[ms]=Buka Tetingkap Baru dengan profil sementara
Name[nb]=??pne et nytt vindu med en midlertidig profil
Name[nl]=Nieuw venster openen met een tijdelijk profiel
Name[pt_BR]=Abrir uma nova janela com um perfil tempor??rio
Name[ro]=Deschide o fereastr?? nou?? cu un profil temporar
Name[ru]=?????????????? ?????????? ???????? ?? ?????????????????? ????????????????
Name[sl]=Odpri novo okno z za??asnim profilom
Name[sv]=??ppna ett nytt f??nster med tempor??r profil
Name[ug]=???????????????? ?????????????? ???????????? ?????????? ???????? ???????????? ??????
Name[vi]=M??? c???a s??? m???i v???i h??? s?? t???m
Name[zh_CN]=????????????????????????????????????
Name[zh_TW]=???????????????????????????????????????
Exec=env BAMF_DESKTOP_FILE_HINT=/var/lib/snapd/desktop/applications/chromium_chromium.desktop /snap/bin/chromium --temp-profile
`

func existsOnMockFileSystem(desktop_file string) (bool, bool, error) {
	existsOnMockFileSystem := strutil.ListContains(mockFileSystem, desktop_file)
	return existsOnMockFileSystem, existsOnMockFileSystem, nil
}

func (s *privilegedDesktopLauncherInternalSuite) mockEnv(key, value string) {
	old := os.Getenv(key)
	os.Setenv(key, value)
	s.AddCleanup(func() {
		os.Setenv(key, old)
	})
}

func (s *privilegedDesktopLauncherInternalSuite) TestDesktopFileSearchPath(c *C) {
	s.mockEnv("HOME", "/home/user")
	s.mockEnv("XDG_DATA_HOME", "")
	s.mockEnv("XDG_DATA_DIRS", "")

	// Default search path
	c.Check(userd.DesktopFileSearchPath(), DeepEquals, []string{
		"/home/user/.local/share/applications",
		"/usr/local/share/applications",
		"/usr/share/applications",
	})

	// XDG_DATA_HOME will override the first path
	s.mockEnv("XDG_DATA_HOME", "/home/user/share")
	c.Check(userd.DesktopFileSearchPath(), DeepEquals, []string{
		"/home/user/share/applications",
		"/usr/local/share/applications",
		"/usr/share/applications",
	})

	// XDG_DATA_DIRS changes the remaining paths
	s.mockEnv("XDG_DATA_DIRS", "/usr/share:/var/lib/snapd/desktop")
	c.Check(userd.DesktopFileSearchPath(), DeepEquals, []string{
		"/home/user/share/applications",
		"/usr/share/applications",
		"/var/lib/snapd/desktop/applications",
	})
}

func (s *privilegedDesktopLauncherInternalSuite) TestDesktopFileIDToFilenameSucceedsWithValidId(c *C) {
	restore := userd.MockRegularFileExists(existsOnMockFileSystem)
	defer restore()
	s.mockEnv("XDG_DATA_DIRS", "/usr/local/share:/usr/share:/var/lib/snapd/desktop")

	var desktopIdTests = []struct {
		id     string
		expect string
	}{
		{"mir-kiosk-scummvm_mir-kiosk-scummvm.desktop", "/var/lib/snapd/desktop/applications/mir-kiosk-scummvm_mir-kiosk-scummvm.desktop"},
		{"foo-bar-baz.desktop", "/var/lib/snapd/desktop/applications/foo-bar/baz.desktop"},
		{"baz-foo-bar.desktop", "/var/lib/snapd/desktop/applications/baz/foo-bar.desktop"},
		{"shadow-test.desktop", "/usr/share/applications/shadow-test.desktop"},
	}

	for _, test := range desktopIdTests {
		actual, err := userd.DesktopFileIDToFilename(test.id)
		c.Assert(err, IsNil)
		c.Assert(actual, Equals, test.expect)
	}
}

func (s *privilegedDesktopLauncherInternalSuite) TestDesktopFileIDToFilenameFailsWithInvalidId(c *C) {
	restore := userd.MockRegularFileExists(existsOnMockFileSystem)
	defer restore()
	s.mockEnv("XDG_DATA_DIRS", "/usr/local/share:/usr/share:/var/lib/snapd/desktop")

	var desktopIdTests = []string{
		"mir-kiosk-scummvm-mir-kiosk-scummvm.desktop",
		"bar-foo-baz.desktop",
		"bar-baz-foo.desktop",
		"foo-bar_foo-bar.desktop",
		// special path segments cannot be smuggled inside desktop IDs
		"bar-..-vlc_vlc.desktop",
		"foo-bar/baz.desktop",
		"bar/../vlc_vlc.desktop",
		"../applications/vlc_vlc.desktop",
		// Other invalid desktop IDs
		"---------foo-bar-baz.desktop",
		"foo-bar-baz.desktop-foo-bar",
		"not-a-dot-desktop",
		"??????????????????????????????-non-ascii-here-too.desktop",
	}

	for _, id := range desktopIdTests {
		_, err := userd.DesktopFileIDToFilename(id)
		c.Check(err, ErrorMatches, `cannot find desktop file for ".*"`, Commentf(id))
	}
}

func (s *privilegedDesktopLauncherInternalSuite) TestVerifyDesktopFileLocation(c *C) {
	restore := userd.MockRegularFileExists(existsOnMockFileSystem)
	defer restore()
	s.mockEnv("XDG_DATA_DIRS", "/usr/local/share:/usr/share:/var/lib/snapd/desktop")

	// Resolved desktop files belonging to snaps will pass verification:
	filename, err := userd.DesktopFileIDToFilename("mir-kiosk-scummvm_mir-kiosk-scummvm.desktop")
	c.Assert(err, IsNil)
	err = userd.VerifyDesktopFileLocation(filename)
	c.Check(err, IsNil)

	// Desktop IDs belonging to host system apps fail:
	filename, err = userd.DesktopFileIDToFilename("shadow-test.desktop")
	c.Assert(err, IsNil)
	err = userd.VerifyDesktopFileLocation(filename)
	c.Check(err, ErrorMatches, "only launching snap applications from /var/lib/snapd/desktop/applications is supported")
}

func (s *privilegedDesktopLauncherInternalSuite) TestParseExecCommandSucceedsWithValidEntry(c *C) {
	var testCases = []struct {
		cmd    string
		expect []string
	}{
		// valid with no exec variables
		{"env BAMF_DESKTOP_FILE_HINT=/var/lib/snapd/desktop/applications/mir-kiosk-scummvm_mir-kiosk-scummvm.desktop /snap/bin/mir-kiosk-scummvm %U",
			[]string{"env", "BAMF_DESKTOP_FILE_HINT=/var/lib/snapd/desktop/applications/mir-kiosk-scummvm_mir-kiosk-scummvm.desktop", "/snap/bin/mir-kiosk-scummvm"}},
		// valid with literal '%' and no exec variables
		{"/snap/bin/foo -f %%bar", []string{"/snap/bin/foo", "-f", "%bar"}},
		{"/snap/bin/foo -f %%bar %%baz", []string{"/snap/bin/foo", "-f", "%bar", "%baz"}},
		// valid where quoted strings are passed through
		{"/snap/bin/foo '-f %U'", []string{"/snap/bin/foo", "-f %U"}},
		{"/snap/bin/foo '-f %%bar'", []string{"/snap/bin/foo", "-f %%bar"}},
		{"/snap/bin/foo '-f %U %%bar'", []string{"/snap/bin/foo", "-f %U %%bar"}},
		{"/snap/bin/foo \"'-f bar'\"", []string{"/snap/bin/foo", "'-f bar'"}},
		{"/snap/bin/foo '\"-f bar\"'", []string{"/snap/bin/foo", "\"-f bar\""}},
		// valid with exec variables stripped out
		{"/snap/bin/foo -f %U", []string{"/snap/bin/foo", "-f"}},
		{"/snap/bin/foo -f %U %i", []string{"/snap/bin/foo", "-f", "--icon", "/snap/chromium/1193/chromium.png"}},
		{"/snap/bin/foo -f %U bar", []string{"/snap/bin/foo", "-f", "bar"}},
		{"/snap/bin/foo -f %U bar %i", []string{"/snap/bin/foo", "-f", "bar", "--icon", "/snap/chromium/1193/chromium.png"}},
		// valid with mixture of literal '%' and exec variables
		{"/snap/bin/foo -f %U %%bar", []string{"/snap/bin/foo", "-f", "%bar"}},
		{"/snap/bin/foo -f %U %i %%bar", []string{"/snap/bin/foo", "-f", "--icon", "/snap/chromium/1193/chromium.png", "%bar"}},
		{"/snap/bin/foo -f %U %%bar %i", []string{"/snap/bin/foo", "-f", "%bar", "--icon", "/snap/chromium/1193/chromium.png"}},
		{"/snap/bin/foo -f %%bar %U %i", []string{"/snap/bin/foo", "-f", "%bar", "--icon", "/snap/chromium/1193/chromium.png"}},
	}

	for _, test := range testCases {
		actual, err := userd.ParseExecCommand(test.cmd, "/snap/chromium/1193/chromium.png")
		comment := Commentf("cmd=%s", test.cmd)
		c.Check(err, IsNil, comment)
		c.Check(actual, DeepEquals, test.expect, comment)
	}
}

func (s *privilegedDesktopLauncherInternalSuite) TestParseExecCommandFailsWithInvalidEntry(c *C) {
	var testCases = []struct {
		cmd string
		err string
	}{
		// Commands may be rejected for bad quoting
		{`/snap/bin/foo "unclosed double quote`, "EOF found when expecting closing quote"},
		{`/snap/bin/foo 'unclosed single quote`, "EOF found when expecting closing quote"},

		// Or use of unexpected unknown variables
		{"/snap/bin/foo %z", `cannot run "/snap/bin/foo %z" due to use of "%z"`},
		{"/snap/bin/foo %", `cannot run "/snap/bin/foo %" due to use of "%"`},
	}

	for _, test := range testCases {
		_, err := userd.ParseExecCommand(test.cmd, "/snap/chromium/1193/chromium.png")
		comment := Commentf("cmd=%s", test.cmd)
		c.Check(err, ErrorMatches, test.err, comment)
	}
}

func (s *privilegedDesktopLauncherInternalSuite) testReadExecCommandFromDesktopFileWithValidContent(c *C, desktopFileContent string) {
	desktopFile := filepath.Join(c.MkDir(), "test.desktop")

	// We need to correct the embedded path to the desktop file before writing the file
	fileContent := strings.Replace(desktopFileContent, "/var/lib/snapd/desktop/applications/chromium_chromium.desktop", desktopFile, -1)
	err := ioutil.WriteFile(desktopFile, []byte(fileContent), 0644)
	c.Assert(err, IsNil)

	exec, icon, err := userd.ReadExecCommandFromDesktopFile(desktopFile)
	c.Assert(err, IsNil)

	c.Check(exec, Equals, fmt.Sprintf("env BAMF_DESKTOP_FILE_HINT=%s %s/chromium %%U", desktopFile, dirs.SnapBinariesDir))
	c.Check(icon, Equals, "/snap/chromium/1193/chromium.png")
}

func (s *privilegedDesktopLauncherInternalSuite) TestReadExecCommandFromDesktopFileWithValidContentPathSnap(c *C) {
	// pick a system known to use /snap/bin for launcher symlinks
	restore := release.MockReleaseInfo(&release.OS{ID: "ubuntu"})
	defer restore()
	dirs.SetRootDir("/")
	defer dirs.SetRootDir("/")
	s.testReadExecCommandFromDesktopFileWithValidContent(c, chromiumDesktopFile)
}

func (s *privilegedDesktopLauncherInternalSuite) TestReadExecCommandFromDesktopFileWithValidContentPathVarLibSnapd(c *C) {
	// pick a system known to use /var/lib/snapd/bin for launcher symlinks
	restore := release.MockReleaseInfo(&release.OS{ID: "fedora"})
	defer restore()
	dirs.SetRootDir("/")
	defer dirs.SetRootDir("/")

	// fix the Exec= line
	fileContentWithVarLibSnapBin := strings.Replace(chromiumDesktopFile, " /snap/bin/chromium", " /var/lib/snapd/snap/bin/chromium", -1)
	s.testReadExecCommandFromDesktopFileWithValidContent(c, fileContentWithVarLibSnapBin)
}

func (s *privilegedDesktopLauncherInternalSuite) TestReadExecCommandFromDesktopFileWithInvalidExec(c *C) {
	desktopFile := filepath.Join(c.MkDir(), "test.desktop")

	err := ioutil.WriteFile(desktopFile, []byte(chromiumDesktopFile), 0644)
	c.Assert(err, IsNil)

	_, _, err = userd.ReadExecCommandFromDesktopFile(desktopFile)
	c.Assert(err, ErrorMatches, `desktop file ".*" has an unsupported 'Exec' value: .*`)
}

func (s *privilegedDesktopLauncherInternalSuite) TestReadExecCommandFromDesktopFileWithNoDesktopEntry(c *C) {
	desktopFile := filepath.Join(c.MkDir(), "test.desktop")

	// We need to correct the embedded path to the desktop file before writing the file
	fileContent := strings.Replace(chromiumDesktopFile, "/var/lib/snapd/desktop/applications/chromium_chromium.desktop", desktopFile, -1)
	fileContent = strings.Replace(fileContent, "[Desktop Entry]", "[garbage]", -1)

	err := ioutil.WriteFile(desktopFile, []byte(fileContent), 0644)
	c.Assert(err, IsNil)

	_, _, err = userd.ReadExecCommandFromDesktopFile(desktopFile)
	c.Assert(err, ErrorMatches, `desktop file ".*" has an unsupported 'Exec' value: ""`)
}

func (s *privilegedDesktopLauncherInternalSuite) TestReadExecCOmmandFromDesktopFileMultipleDesktopEntrySections(c *C) {
	desktopFile := filepath.Join(c.MkDir(), "test.desktop")
	c.Assert(ioutil.WriteFile(desktopFile, []byte(`[Desktop Entry]
Exec=foo

[Desktop Entry]
Exec=bar
`), 0644), IsNil)

	_, _, err := userd.ReadExecCommandFromDesktopFile(desktopFile)
	c.Check(err, ErrorMatches, `desktop file ".*" has multiple \[Desktop Entry\] sections`)
}

func (s *privilegedDesktopLauncherInternalSuite) TestReadExecCommandFromDesktopFileWithNoFile(c *C) {
	desktopFile := filepath.Join(c.MkDir(), "test.desktop")

	_, _, err := userd.ReadExecCommandFromDesktopFile(desktopFile)
	c.Assert(err, ErrorMatches, `open .*: no such file or directory`)
}
