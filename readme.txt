
�P�Dgolang �E�B���h�E�Y�ł��C���X�g�[��

https://golang.org/dl/


�Q�Dgoenv.bat �����s

�@GOROOT �� golang ���C���X�g�[�������f�B���N�g����ݒ�
�@GOPATH �� ���̂ł��ꂤ�Ƃ��ݒ�
�@
  goenv.bat �� set GOROOT=, set GOPATH= �����������Ď��s���Ă��������B
  
�R�D�ݒ�̏C��

�@���x�̏�����ύX����ꍇ�́A
�@service.go ��
  const (
    // �x�����x
    warningLimit  = 33.5
    // �G���[���x
    emergeLimit   = 35.0
  )
�@���C�����܂��B
�@
�@���[������ύX����ꍇ��
�@sendmail.go
�@�֐� sendMail �̒萔�����������܂�
  func sendMail(name string, title string, description string) error {
    from     := "daemon@foo dot com"
    to       := "tweet@foo dot com"
    user     := "smtp user@foo dot com"
    password := "your password"
    subject  := title
    server   := "foo dot com"
    port     := "587"
�@

�S�Dbuild.bat �����s

  HnsTemp.exe ����������܂��B
  
�T�D�C���X�g�[��

  HnsTemp.exe �� USBMeter.DLL ���R�s�[����
  
  HnsTemp install �����s���܂��B
  �T�[�r�X�Ƃ��ēo�^����܂��B
  HnsTemp�T�[�r�X���J���Ď������s�ɕύX���܂��B
  
  HnsTemp remove �ŃT�[�r�X����폜����܂��B
  


�⑫�F

USBMeter.dll �́A���L�̃L�b�g�ł��B
http://strawberry-linux.com/catalog/items?code=52001