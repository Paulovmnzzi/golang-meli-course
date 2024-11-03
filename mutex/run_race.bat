@echo off
setlocal enabledelayedexpansion

set count=100
set cosmos_count=0
set universe_count=0

for /l %%i in (1, 1, %count%) do (
    echo Run #%%i
    for /f "tokens=*" %%a in ('go run -race .') do (
        echo %%a
        if "%%a"=="Hello cosmos!" (
            set /a cosmos_count+=1
        ) else if "%%a"=="Hello universe!" (
            set /a universe_count+=1
        )
    )
    echo -------------------------
)

echo Total runs: %count%
echo "Hello cosmos!" count: %cosmos_count%
echo "Hello universe!" count: %universe_count%

endlocal