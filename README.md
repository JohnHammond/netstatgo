# netstatgo

> John Hammond | Monday, April 10, 2023

---------------

Crappy Golang code to list local listening ports and their associated processes, much like `netstat -anob` on Windows.

To receive all possible information, run as an administrator.

## Sample Output

```
Proto   Local Address             Foreign Address           State      PID/Name
TCP     0.0.0.0:135               0.0.0.0:0                 LISTEN     1948      /svchost.exe
TCP     0.0.0.0:445               0.0.0.0:0                 LISTEN     4         /System
TCP     0.0.0.0:903               0.0.0.0:0                 LISTEN     6092      /vmware-authd.exe
TCP     0.0.0.0:913               0.0.0.0:0                 LISTEN     6092      /vmware-authd.exe
TCP     0.0.0.0:1824              0.0.0.0:0                 LISTEN     18144     /WaveLink.exe
TCP     0.0.0.0:3389              0.0.0.0:0                 LISTEN     1220      /svchost.exe
TCP     0.0.0.0:5040              0.0.0.0:0                 LISTEN     11008     /svchost.exe
TCP     0.0.0.0:7680              0.0.0.0:0                 LISTEN     10192     /svchost.exe
TCP     0.0.0.0:49664             0.0.0.0:0                 LISTEN     1632      /lsass.exe
TCP     0.0.0.0:49665             0.0.0.0:0                 LISTEN     1536      /wininit.exe
TCP     0.0.0.0:49666             0.0.0.0:0                 LISTEN     1052      /svchost.exe
TCP     0.0.0.0:49667             0.0.0.0:0                 LISTEN     2908      /svchost.exe
TCP     0.0.0.0:49668             0.0.0.0:0                 LISTEN     3104      /svchost.exe
TCP     0.0.0.0:49669             0.0.0.0:0                 LISTEN     4884      /spoolsv.exe
TCP     0.0.0.0:49671             0.0.0.0:0                 LISTEN     1612      /services.exe
TCP     127.0.0.1:28198           0.0.0.0:0                 LISTEN     18144     /WaveLink.exe
TCP     127.0.0.1:37373           0.0.0.0:0                 LISTEN     32188     /gocode.exe
TCP     127.0.0.1:49670           0.0.0.0:0                 LISTEN     5512      /dirmngr.exe
TCP     127.0.0.1:50796           0.0.0.0:0                 LISTEN     5212      /dbus-daemon.exe
TCP     127.0.0.1:50796           0.0.0.0:0                 LISTEN     5212      /dbus-daemon.exe
TCP     127.0.0.1:56993           127.0.0.1:56994           ESTABLISHED 27928     /Zoom.exe
TCP     127.0.0.1:56994           127.0.0.1:56993           ESTABLISHED 27928     /Zoom.exe
TCP     127.0.0.1:58582           127.0.0.1:58583           ESTABLISHED 11252     /vmware.exe
TCP     127.0.0.1:58583           127.0.0.1:58582           ESTABLISHED 11252     /vmware.exe
TCP     127.0.0.1:58586           127.0.0.1:58587           ESTABLISHED 11252     /vmware.exe
TCP     127.0.0.1:58587           127.0.0.1:58586           ESTABLISHED 11252     /vmware.exe
TCP     127.0.0.1:58588           127.0.0.1:58589           ESTABLISHED 11252     /vmware.exe
TCP     127.0.0.1:58589           127.0.0.1:58588           ESTABLISHED 11252     /vmware.exe
```