package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"testing"

	check "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type FSSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { check.TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = check.Suite(&FSSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *FSSuite) TestList(c *check.C) {
	tmpDir := c.MkDir()

	os.Mkdir(tmpDir+"/.dir0", 0755)

	os.Create(tmpDir + "/.file0")

	ioutil.WriteFile(tmpDir+"/file1.mp3", []byte("TESTDATA12345678"), 644)
	ioutil.WriteFile(tmpDir+"/file2.jpg", []byte("TESTDATA"), 644)

	os.Mkdir(tmpDir+"/dir1", 0755)
	os.Mkdir(tmpDir+"/dir2", 0755)

	os.Create(tmpDir + "/dir1/file3.mp3")
	os.Create(tmpDir + "/dir2/file4.wav")

	os.Mkdir(tmpDir+"/dir1/dir3", 0755)

	listing1 := List(tmpDir, false)
	listing2 := List(tmpDir, true)
	listing3 := ListAll(tmpDir, false)
	listing4 := ListAll(tmpDir, true, ListingFilter{})
	listing5 := ListAllDirs(tmpDir, false)
	listing6 := ListAllDirs(tmpDir, true, ListingFilter{})
	listing7 := ListAllFiles(tmpDir, false)
	listing8 := ListAllFiles(tmpDir, true)
	listing9 := ListAllFiles(tmpDir, true, ListingFilter{MatchPatterns: []string{"*.mp3", "*.wav"}})
	listing10 := ListAllFiles(tmpDir, true, ListingFilter{NotMatchPatterns: []string{"*.mp3"}})
	listing11 := List(tmpDir, true, ListingFilter{Perms: "DR"})
	listing12 := List(tmpDir, true, ListingFilter{NotPerms: "DR"})
	listing13 := ListAllFiles(tmpDir, true, ListingFilter{NotMatchPatterns: []string{"*.mp3"}, SizeZero: true})
	listing14 := ListAllFiles(tmpDir, false, ListingFilter{SizeEqual: 16})
	listing15 := ListAllFiles(tmpDir, false, ListingFilter{SizeLess: 12, SizeGreater: 5})
	listing16 := ListAllFiles(tmpDir, false, ListingFilter{SizeGreater: 12})
	listing17 := List(
		tmpDir, false,
		ListingFilter{
			ATimeOlder:   2524608000,
			CTimeOlder:   2524608000,
			MTimeOlder:   2524608000,
			ATimeYounger: 1,
			CTimeYounger: 1,
			MTimeYounger: 1,
		},
	)

	sort.Strings(listing1)
	sort.Strings(listing2)
	sort.Strings(listing3)
	sort.Strings(listing4)
	sort.Strings(listing5)
	sort.Strings(listing6)
	sort.Strings(listing7)
	sort.Strings(listing8)
	sort.Strings(listing9)
	sort.Strings(listing10)
	sort.Strings(listing11)
	sort.Strings(listing12)
	sort.Strings(listing13)
	sort.Strings(listing14)
	sort.Strings(listing15)
	sort.Strings(listing16)
	sort.Strings(listing17)

	c.Assert(
		listing1,
		check.DeepEquals,
		[]string{".dir0", ".file0", "dir1", "dir2", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing2,
		check.DeepEquals,
		[]string{"dir1", "dir2", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing3,
		check.DeepEquals,
		[]string{".dir0", ".file0", "dir1", "dir1/dir3", "dir1/file3.mp3", "dir2", "dir2/file4.wav", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing4,
		check.DeepEquals,
		[]string{"dir1", "dir1/dir3", "dir1/file3.mp3", "dir2", "dir2/file4.wav", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing5,
		check.DeepEquals,
		[]string{".dir0", "dir1", "dir1/dir3", "dir2"},
	)

	c.Assert(
		listing6,
		check.DeepEquals,
		[]string{"dir1", "dir1/dir3", "dir2"},
	)

	c.Assert(
		listing7,
		check.DeepEquals,
		[]string{".file0", "dir1/file3.mp3", "dir2/file4.wav", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing8,
		check.DeepEquals,
		[]string{"dir1/file3.mp3", "dir2/file4.wav", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing9,
		check.DeepEquals,
		[]string{"dir1/file3.mp3", "dir2/file4.wav", "file1.mp3"},
	)

	c.Assert(
		listing10,
		check.DeepEquals,
		[]string{"dir2/file4.wav", "file2.jpg"},
	)

	c.Assert(
		listing11,
		check.DeepEquals,
		[]string{"dir1", "dir2"},
	)

	c.Assert(
		listing12,
		check.DeepEquals,
		[]string{"file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing13,
		check.DeepEquals,
		[]string{"dir2/file4.wav"},
	)

	c.Assert(
		listing14,
		check.DeepEquals,
		[]string{"file1.mp3"},
	)

	c.Assert(
		listing15,
		check.DeepEquals,
		[]string{"file2.jpg"},
	)

	c.Assert(
		listing16,
		check.DeepEquals,
		[]string{"file1.mp3"},
	)

	c.Assert(
		listing17,
		check.DeepEquals,
		[]string{".dir0", ".file0", "dir1", "dir2", "file1.mp3", "file2.jpg"},
	)

	c.Assert(readDir("/not_exist"), check.IsNil)

	c.Assert(ListingFilter{ATimeOlder: 1}.hasTimes(), check.Equals, true)
	c.Assert(ListingFilter{ATimeYounger: 1}.hasTimes(), check.Equals, true)
	c.Assert(ListingFilter{CTimeOlder: 1}.hasTimes(), check.Equals, true)
	c.Assert(ListingFilter{CTimeYounger: 1}.hasTimes(), check.Equals, true)
	c.Assert(ListingFilter{MTimeOlder: 1}.hasTimes(), check.Equals, true)
	c.Assert(ListingFilter{MTimeYounger: 1}.hasTimes(), check.Equals, true)
}

func (s *FSSuite) TestListToAbsolute(c *check.C) {
	list := []string{"1", "2", "3"}

	ListToAbsolute("A", list)

	c.Assert(list, check.DeepEquals, []string{"A/1", "A/2", "A/3"})
}

func (s *FSSuite) TestProperPath(c *check.C) {
	tmpFile := c.MkDir() + "/test.txt"

	os.OpenFile(tmpFile, os.O_CREATE, 0644)

	paths := []string{"/etc/sudoers", "/etc/passwd", tmpFile}

	c.Assert(ProperPath("DR", paths), check.Equals, "")
	c.Assert(ProperPath("FR", paths), check.Equals, "/etc/passwd")
	c.Assert(ProperPath("FRW", paths), check.Equals, tmpFile)
	c.Assert(ProperPath("FRWS", paths), check.Equals, "")
	c.Assert(ProperPath("F", paths), check.Equals, "/etc/sudoers")

	os.Remove(tmpFile)
}

func (s *FSSuite) TestWalker(c *check.C) {
	tmpDir := c.MkDir()

	os.Chdir(tmpDir)

	tmpDir, _ = os.Getwd()

	os.MkdirAll(tmpDir+"/dir1/dir2/dir3/dir4", 0755)
	os.Chdir(tmpDir)

	c.Assert(Current(), check.Equals, tmpDir)
	c.Assert(Pop(), check.Equals, tmpDir)

	dirStack = nil

	c.Assert(Push("dir1"), check.Equals, tmpDir+"/dir1")
	c.Assert(Push("dir9"), check.Equals, "")
	c.Assert(Push("dir2/dir3"), check.Equals, tmpDir+"/dir1/dir2/dir3")
	c.Assert(Push("dir4"), check.Equals, tmpDir+"/dir1/dir2/dir3/dir4")
	c.Assert(Push("dir9"), check.Equals, "")
	c.Assert(Pop(), check.Equals, tmpDir+"/dir1/dir2/dir3")
	c.Assert(Pop(), check.Equals, tmpDir+"/dir1")
	c.Assert(Pop(), check.Equals, tmpDir)
	c.Assert(Pop(), check.Equals, tmpDir)

	c.Assert(Push("dir1"), check.Equals, tmpDir+"/dir1")
	c.Assert(Push("dir2"), check.Equals, tmpDir+"/dir1/dir2")
	c.Assert(Push("dir3"), check.Equals, tmpDir+"/dir1/dir2/dir3")
	os.RemoveAll(tmpDir + "/dir1/dir2")
	c.Assert(Pop(), check.Equals, "")
}

func (s *FSSuite) TestGetSize(c *check.C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/test.file"

	if ioutil.WriteFile(tmpFile, []byte("TEST\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	c.Assert(GetSize(""), check.Equals, int64(-1))
	c.Assert(GetSize("/not_exist"), check.Equals, int64(-1))
	c.Assert(GetSize(tmpFile), check.Equals, int64(5))
}

func (s *FSSuite) TestGetTime(c *check.C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/test.file"

	if ioutil.WriteFile(tmpFile, []byte("TEST\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	at, mt, ct, err := GetTimes(tmpFile)

	c.Assert(err, check.IsNil)
	c.Assert(at.IsZero(), check.Equals, false)
	c.Assert(mt.IsZero(), check.Equals, false)
	c.Assert(ct.IsZero(), check.Equals, false)

	at, mt, ct, err = GetTimes("")

	c.Assert(err, check.NotNil)
	c.Assert(err, check.Equals, ErrEmptyPath)
	c.Assert(at.IsZero(), check.Equals, true)
	c.Assert(mt.IsZero(), check.Equals, true)
	c.Assert(ct.IsZero(), check.Equals, true)

	at, mt, ct, err = GetTimes("/not_exist")

	c.Assert(err, check.NotNil)
	c.Assert(at.IsZero(), check.Equals, true)
	c.Assert(mt.IsZero(), check.Equals, true)
	c.Assert(ct.IsZero(), check.Equals, true)

	ats, mts, cts, err := GetTimestamps(tmpFile)

	c.Assert(err, check.IsNil)
	c.Assert(ats, check.Not(check.Equals), int64(-1))
	c.Assert(mts, check.Not(check.Equals), int64(-1))
	c.Assert(cts, check.Not(check.Equals), int64(-1))

	ats, mts, cts, err = GetTimestamps("")

	c.Assert(err, check.NotNil)
	c.Assert(err, check.Equals, ErrEmptyPath)
	c.Assert(ats, check.Equals, int64(-1))
	c.Assert(mts, check.Equals, int64(-1))
	c.Assert(cts, check.Equals, int64(-1))

	ats, mts, cts, err = GetTimestamps("/not_exist")

	c.Assert(err, check.NotNil)
	c.Assert(ats, check.Equals, int64(-1))
	c.Assert(mts, check.Equals, int64(-1))
	c.Assert(cts, check.Equals, int64(-1))

	at, err = GetATime(tmpFile)

	c.Assert(err, check.IsNil)
	c.Assert(at.IsZero(), check.Equals, false)

	mt, err = GetMTime(tmpFile)

	c.Assert(err, check.IsNil)
	c.Assert(mt.IsZero(), check.Equals, false)

	ct, err = GetCTime(tmpFile)

	c.Assert(err, check.IsNil)
	c.Assert(ct.IsZero(), check.Equals, false)

	at, err = GetATime("")

	c.Assert(err, check.NotNil)
	c.Assert(at.IsZero(), check.Equals, true)

	mt, err = GetMTime("")

	c.Assert(err, check.NotNil)
	c.Assert(mt.IsZero(), check.Equals, true)

	ct, err = GetCTime("")

	c.Assert(err, check.NotNil)
	c.Assert(ct.IsZero(), check.Equals, true)
}

func (s *FSSuite) TestGetOwner(c *check.C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/test.file"

	if ioutil.WriteFile(tmpFile, []byte("TEST\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	uid, gid, err := GetOwner(tmpFile)

	c.Assert(err, check.IsNil)
	c.Assert(uid, check.Not(check.Equals), -1)
	c.Assert(gid, check.Not(check.Equals), -1)

	uid, gid, err = GetOwner("")

	c.Assert(err, check.NotNil)
	c.Assert(err, check.Equals, ErrEmptyPath)
	c.Assert(uid, check.Equals, -1)
	c.Assert(gid, check.Equals, -1)

	uid, gid, err = GetOwner("/not_exist")

	c.Assert(err, check.NotNil)
	c.Assert(uid, check.Equals, -1)
	c.Assert(gid, check.Equals, -1)
}

func (s *FSSuite) TestIsEmptyDir(c *check.C) {
	tmpDir1 := c.MkDir()
	tmpDir2 := c.MkDir()
	tmpFile := tmpDir1 + "/test.file"

	if ioutil.WriteFile(tmpFile, []byte("TEST\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	c.Assert(IsEmptyDir(tmpDir1), check.Equals, false)
	c.Assert(IsEmptyDir(tmpDir2), check.Equals, true)
	c.Assert(IsEmptyDir(""), check.Equals, false)
	c.Assert(IsEmptyDir("/not_exist"), check.Equals, false)
}

func (s *FSSuite) TestIsNonEmpty(c *check.C) {
	tmpDir := c.MkDir()
	tmpFile1 := tmpDir + "/test1.file"
	tmpFile2 := tmpDir + "/test2.file"

	if ioutil.WriteFile(tmpFile1, []byte("TEST\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	if ioutil.WriteFile(tmpFile2, []byte(""), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	c.Assert(IsNonEmpty(""), check.Equals, false)
	c.Assert(IsNonEmpty("/not_exist"), check.Equals, false)
	c.Assert(IsNonEmpty(tmpFile2), check.Equals, false)
	c.Assert(IsNonEmpty(tmpFile1), check.Equals, true)
}

func (s *FSSuite) TestTypeChecks(c *check.C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/test.file"
	tmpLink := tmpDir + "/test.link"

	if ioutil.WriteFile(tmpFile, []byte("TEST\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	if os.Symlink("123", tmpLink) != nil {
		c.Fatal("Can't create link")
	}

	c.Assert(IsExist(""), check.Equals, false)
	c.Assert(IsExist("/not_exist"), check.Equals, false)
	c.Assert(IsExist(tmpFile), check.Equals, true)

	c.Assert(IsRegular(""), check.Equals, false)
	c.Assert(IsRegular("/not_exist"), check.Equals, false)
	c.Assert(IsRegular(tmpFile), check.Equals, true)
	c.Assert(IsRegular(tmpLink), check.Equals, false)

	c.Assert(IsLink(""), check.Equals, false)
	c.Assert(IsLink("/not_exist"), check.Equals, false)
	c.Assert(IsLink(tmpFile), check.Equals, false)
	c.Assert(IsLink(tmpLink), check.Equals, true)

	c.Assert(IsCharacterDevice(""), check.Equals, false)
	c.Assert(IsCharacterDevice("/not_exist"), check.Equals, false)
	c.Assert(IsCharacterDevice(tmpFile), check.Equals, false)
	c.Assert(IsCharacterDevice("/dev/tty"), check.Equals, true)

	c.Assert(IsBlockDevice(""), check.Equals, false)
	c.Assert(IsBlockDevice("/not_exist"), check.Equals, false)
	c.Assert(IsBlockDevice(tmpFile), check.Equals, false)

	switch {
	case IsExist("/dev/sda"):
		c.Assert(IsBlockDevice("/dev/sda"), check.Equals, true)
	case IsExist("/dev/vda"):
		c.Assert(IsBlockDevice("/dev/vda"), check.Equals, true)
	case IsExist("/dev/hda"):
		c.Assert(IsBlockDevice("/dev/hda"), check.Equals, true)
	case IsExist("/dev/disk0"):
		c.Assert(IsBlockDevice("/dev/disk0"), check.Equals, true)
	}

	c.Assert(IsDir(""), check.Equals, false)
	c.Assert(IsDir("/not_exist"), check.Equals, false)
	c.Assert(IsDir(tmpFile), check.Equals, false)
	c.Assert(IsDir(tmpDir), check.Equals, true)

	c.Assert(IsSocket(""), check.Equals, false)
	c.Assert(IsSocket("/not_exist"), check.Equals, false)
	c.Assert(IsSocket(tmpFile), check.Equals, false)
	c.Assert(IsSocket(tmpDir), check.Equals, false)

	switch {
	case IsExist("/var/run/mDNSResponder"):
		c.Assert(IsSocket("/var/run/mDNSResponder"), check.Equals, true)
	case IsExist("/dev/log"):
		c.Assert(IsSocket("/dev/log"), check.Equals, true)
	}
}

func (s *FSSuite) TestPermChecks(c *check.C) {
	tmpDir := c.MkDir()
	tmpFile1 := tmpDir + "/test1.file"
	tmpFile2 := tmpDir + "/test2.file"
	tmpFile3 := tmpDir + "/test3.file"
	tmpFile4 := tmpDir + "/test4.file"
	tmpFile5 := tmpDir + "/test5.file"
	tmpFile6 := tmpDir + "/test6.file"
	tmpFile7 := tmpDir + "/test7.file"
	tmpFile8 := tmpDir + "/test8.file"
	tmpFile9 := tmpDir + "/test9.file"

	for i := 1; i <= 9; i++ {
		if ioutil.WriteFile(fmt.Sprintf("%s/test%d.file", tmpDir, i), []byte(""), 0644) != nil {
			c.Fatal("Can't create temporary file")
		}
	}

	os.Chmod(tmpFile1, 0400)
	os.Chmod(tmpFile2, 0040)
	os.Chmod(tmpFile3, 0004)
	os.Chmod(tmpFile4, 0200)
	os.Chmod(tmpFile5, 0020)
	os.Chmod(tmpFile6, 0002)
	os.Chmod(tmpFile7, 0100)
	os.Chmod(tmpFile8, 0010)
	os.Chmod(tmpFile9, 0001)

	c.Assert(IsReadable(""), check.Equals, false)
	c.Assert(IsReadable("/not_exist"), check.Equals, false)
	c.Assert(IsReadable(tmpFile1), check.Equals, true)
	c.Assert(IsReadable(tmpFile2), check.Equals, true)
	c.Assert(IsReadable(tmpFile3), check.Equals, true)

	c.Assert(IsWritable(""), check.Equals, false)
	c.Assert(IsWritable("/not_exist"), check.Equals, false)
	c.Assert(IsWritable(tmpFile4), check.Equals, true)
	c.Assert(IsWritable(tmpFile5), check.Equals, true)
	c.Assert(IsWritable(tmpFile6), check.Equals, true)

	c.Assert(IsExecutable(""), check.Equals, false)
	c.Assert(IsExecutable("/not_exist"), check.Equals, false)
	c.Assert(IsExecutable(tmpFile7), check.Equals, true)
	c.Assert(IsExecutable(tmpFile8), check.Equals, true)
	c.Assert(IsExecutable(tmpFile9), check.Equals, true)
	c.Assert(IsExecutable(tmpFile1), check.Equals, false)
}

func (s *FSSuite) TestCheckPerms(c *check.C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/test.file"
	tmpLink := tmpDir + "/test.link"

	if ioutil.WriteFile(tmpFile, []byte(""), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	if os.Symlink("123", tmpLink) != nil {
		c.Fatal("Can't create link")
	}

	c.Assert(CheckPerms("", tmpFile), check.Equals, false)
	c.Assert(CheckPerms("FR", ""), check.Equals, false)
	c.Assert(CheckPerms("FR", "/not_exist"), check.Equals, false)

	c.Assert(CheckPerms("F", tmpDir), check.Equals, false)
	c.Assert(CheckPerms("D", tmpFile), check.Equals, false)
	c.Assert(CheckPerms("L", tmpFile), check.Equals, false)
	c.Assert(CheckPerms("X", tmpFile), check.Equals, false)
	c.Assert(CheckPerms("S", tmpFile), check.Equals, false)

	c.Assert(CheckPerms("W", tmpFile), check.Equals, true)
	c.Assert(CheckPerms("R", tmpFile), check.Equals, true)
}

func (s *FSSuite) TestGetPerms(c *check.C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/test.file"

	if ioutil.WriteFile(tmpFile, []byte("TEST\n"), 0764) != nil {
		c.Fatal("Can't create temporary file")
	}

	os.Chmod(tmpFile, 0764)

	c.Assert(GetPerms(""), check.Equals, os.FileMode(0))
	c.Assert(GetPerms(tmpFile), check.Equals, os.FileMode(0764))
}

func (s *FSSuite) TestLineCount(c *check.C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/test.file"

	if ioutil.WriteFile(tmpFile, []byte("1\n2\n3\n4\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	c.Assert(LineCount(""), check.Equals, -1)
	c.Assert(LineCount("/not_exist"), check.Equals, -1)
	c.Assert(LineCount(tmpFile), check.Equals, 4)
}

func (s *FSSuite) TestCopyFile(c *check.C) {
	tmpDir1 := c.MkDir()
	tmpDir2 := c.MkDir()
	tmpDir3 := c.MkDir()
	tmpFile1 := tmpDir1 + "/test1.file"
	tmpFile2 := tmpDir2 + "/test2.file"
	tmpFile3 := tmpDir1 + "/test3.file"

	if ioutil.WriteFile(tmpFile1, []byte("TEST\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	if ioutil.WriteFile(tmpFile2, []byte("TEST1234TEST\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	if ioutil.WriteFile(tmpFile3, []byte(""), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	os.Chmod(tmpFile3, 0111)
	os.Chmod(tmpDir3, 0500)

	c.Assert(CopyFile("", tmpFile2), check.NotNil)
	c.Assert(CopyFile(tmpFile1, ""), check.NotNil)
	c.Assert(CopyFile("/not_exist", tmpFile2), check.NotNil)
	c.Assert(CopyFile(tmpDir1, tmpFile2), check.NotNil)
	c.Assert(CopyFile(tmpFile3, tmpFile2), check.NotNil)
	c.Assert(CopyFile(tmpFile1, "/not_exist/test.file"), check.NotNil)
	c.Assert(CopyFile(tmpFile1, tmpDir3+"/test.file"), check.NotNil)
	c.Assert(CopyFile(tmpFile1, tmpDir2), check.NotNil)
	c.Assert(CopyFile(tmpFile1, tmpFile3), check.NotNil)

	c.Assert(CopyFile(tmpFile1, tmpFile2, 0600), check.IsNil)
	c.Assert(GetSize(tmpFile2), check.Equals, int64(5))
	c.Assert(GetPerms(tmpFile2), check.Equals, os.FileMode(0600))

	os.Remove(tmpFile2)

	c.Assert(CopyFile(tmpFile1, tmpFile2, 0600), check.IsNil)
	c.Assert(GetSize(tmpFile2), check.Equals, int64(5))
	c.Assert(GetPerms(tmpFile2), check.Equals, os.FileMode(0600))

	_disableCopyFileChecks = true

	c.Assert(CopyFile("", tmpFile2), check.NotNil)
	c.Assert(CopyFile(tmpFile1, ""), check.NotNil)
}

func (s *FSSuite) TestMoveFile(c *check.C) {
	tmpDir := c.MkDir()
	tmpDir2 := c.MkDir()
	tmpFile1 := tmpDir + "/test1.file"
	tmpFile2 := tmpDir + "/test2.file"
	tmpFile3 := tmpDir + "/test3.file"

	if ioutil.WriteFile(tmpFile1, []byte("TEST\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	if ioutil.WriteFile(tmpFile3, []byte("TEST\n"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	os.Chmod(tmpFile3, 0111)
	os.Chmod(tmpDir2, 0500)

	c.Assert(MoveFile("", tmpFile2), check.NotNil)
	c.Assert(MoveFile(tmpFile1, ""), check.NotNil)
	c.Assert(MoveFile("/not_exist", tmpFile2), check.NotNil)
	c.Assert(MoveFile(tmpDir, tmpFile2), check.NotNil)
	c.Assert(MoveFile(tmpFile3, tmpFile2), check.NotNil)
	c.Assert(MoveFile(tmpFile1, "/not_exist/file.test"), check.NotNil)
	c.Assert(MoveFile(tmpFile1, tmpDir2+"/file.test"), check.NotNil)

	c.Assert(MoveFile(tmpFile1, tmpFile2), check.IsNil)
	c.Assert(MoveFile(tmpFile2, tmpFile1, 0600), check.IsNil)

	_disableMoveFileChecks = true

	c.Assert(MoveFile("", tmpFile2), check.NotNil)
}

func (s *FSSuite) TestInternal(c *check.C) {
	c.Assert(getGIDList(nil), check.IsNil)

	c.Assert(isReadableStat(nil, 0, nil), check.Equals, true)
	c.Assert(isWritableStat(nil, 0, nil), check.Equals, true)
	c.Assert(isExecutableStat(nil, 0, nil), check.Equals, true)

	n, _ := fixCount(-100, nil)

	c.Assert(n, check.Equals, 0)
}
