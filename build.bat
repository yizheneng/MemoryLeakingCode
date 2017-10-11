cd others
qtdeploy.exe build
cd ../src/
qtdeploy.exe build
cd ../
copy ./src/deploy/windows/src.exe ./MonitoringCenter.exe
MonitoringCenter.exe