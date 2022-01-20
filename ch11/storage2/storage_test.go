package storage

import (
  "strings"
  "testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
  // もとのnotifyUserを保存しておいて回復する
  saved := notifyUser
  defer func() { notifyUser = saved}()

  var notifiedUser, notifiedMsg string
  notifyUser = func(user, msg string) {
    notifiedUser, notifiedMsg = user, msg
  }

  // 980MBが使われた状態を装う

  const user = "joe@example.com"
  CheckQuota(user)
  if notifiedUser == "" && notifiedMsg == "" {
    t.Fatalf("notifyUser not called")
  }
  if notifiedUser != user {
    t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
  }
  const wantSubstring = "98% of your quota"
  if !strings.Contains(notifiedMsg, wantSubstring) {
    t.Errorf("unexpected notification message <<%s.>, "+
      "want substring %q", notifiedMsg, wantSubstring)
  }
}

