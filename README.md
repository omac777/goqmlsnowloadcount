So here is the entire recipe from start to finish <br>
to get a go-qml based skeleton application running on Google Nexus 4 with Ubuntu Touch.<br>
In this example we will use the recipe to demonstrate <br>
running goqmlsnowloadcount, an Ubuntu Touch wizard-gui <br>
with next/previous/finish buttons <br>
that uses Ubuntu touch tab components. <br>
The golang backend in goqmlsnowloadcount currently does nothing <br>
except generate log messages as a proof of concept.<br>
<br>
Step 1)(15-20 minutes)<br>
Flash your Google Nexus 4 phone with the latest android firmware.<br>
This is to get the radio firmware from the latest:<br>
https://dl.google.com/dl/android/aosp/occam-kot49h-factory-02e344de.tgz<br>
tar xvf occam-kot49h-factory-02e344de.tgz<br>
Connect your phone to your computer.<br>
cd occam-kot49h/<br>
The following line puts your phone into recovery mode:<br>
adb reboot-bootloader<br>
Then this command flashes the latest android firmware:<br>
./flash-all.sh<br>
<br>
Step 2)(15-20 minutes)<br>
Flash your Google Nexus 4 phone with the latest version of Ubuntu Touch devel-propose r301 version.<br>
Make sure your phone is connected to the computer.<br>
Make sure to put the phone in recovery mode:<br>
adb reboot-bootloader<br>
sudo apt-get install ubuntu-device-flash<br>
ubuntu-device-flash --revision=301 --channel=devel-proposed --bootstrap<br>
<br>
Step 3)(1 minute)<br>
Start talking to your phone from your computer<br>
On your computer, place this in a file called /home/youruser/.bash_aliases<br>
alias sd='adb shell start ssh; \<br>
    adb forward tcp:2222 tcp:22; \<br>
    ssh-keygen -f ~/.ssh/known_hosts -R [localhost]:2222; \<br>
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no \<br>
<br>
Now open up a ssh session to your phone:<br>
sd<br>
It will return with the following prompt:<br>
phablet@ubuntu-phablet:~/$<br>
<br>
Step 3)(5 minutes)<br>
From your ssh session, install the prerequisites for running go-qml infrastructure.<br>
mkdir -p /home/phablet/goqmlstuff<br>
export GOPATH=/home/phablet/goqmlstuff<br>
sudo apt-get install git golang pkg-config golang-go.tools<br>
sudo apt-get install qtbase5-private-dev<br>
sudo apt-get install qtdeclarative5-private-dev<br>
sudo apt-get install libqt5opengl5-dev<br>
sudo apt-get install gccgo gccgo-go<br>
Once you have install the packages, you don't need the .deb files taking up precious space <br>
so we remove them to avoid getting "storage exceeded" errors.<br>
sudo apt-get clean<br>
sudo apt-get autoremove<br>
<br>
These packages are huge so install them and then remove the .deb files.<br>
sudo apt-get install g++ g++-multilib<br>
sudo apt-get clean<br>
sudo apt-get autoremove<br>
<br>
Step 4)(2 minutes)<br>
Install go-qml<br>
go get gopkg.in/qml.v0<br>
cd /home/phablet/goqmlstuff/src/gopkg.in/qml.v0/examples/customtype<br>
go build<br>
This should prove that you can build golang qml apps.<br>
<br>
Step 5)(2 seconds)<br>
Install goqmlsnowloadcount<br>
cd /home/phablet/goqmlstuff<br>
go get github.com/omac777/goqmlsnowloadcount
<br>
Step 6)(3 seconds)<br>
Build goqmlsnowloadcount<br>
cd /home/phablet/goqmlstuff/src/github.com/omac777/goqmlsnowloadcount<br>
go build<br>
<br>
Step 7)(2 seconds)<br>
Run goqmlsnowloadcount<br>
cd /home/phablet/.local/share/applications<br>
wget https://niemeyer.s3.amazonaws.com/test.desktop<br>
APP_ID=goqmlsnowloadcount ./goqmlsnowloadcount --desktop_file_hint=/home/phablet/.local/share/applications/test.desktop<br>
<br>
Step 8)<br>
Debugging the application directly on the Google Nexus 4<br>
8.1)Start the GNU Debugger directly on the Google Nexus 4:<br>
APP_ID=goqmlsnowloadcount gdb ./goqmlsnowloadcount<br>
8.2)From within gdb, run the application using the required desktop_file_hint argument:<br>
r --desktop_file_hint=/home/phablet/.local/share/applications/test.desktop<br>
<br>
That's it.<br>
Have fun coding/debugging your own stuff with golang and qml directly on your Google Nexus 4 :)<br>
