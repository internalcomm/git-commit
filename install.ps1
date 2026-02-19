$ErrorActionPreference = "Stop"

$App        = "git-commit"
$BaseUrl    = "https://github.com/internalcomm/git-commit/releases/latest/download"
$ZipName    = "git-commit-dist.zip"
$TmpDir     = "$env:TEMP\$App-install"
$InstallDir = "$env:USERPROFILE\bin"
$Target     = "$InstallDir\$App.exe"

[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

New-Item -ItemType Directory -Path $TmpDir     -Force | Out-Null
New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null

Write-Host "Downloading $App..." -ForegroundColor Cyan
try {
    Invoke-WebRequest "$BaseUrl/$ZipName" -OutFile "$TmpDir\$ZipName" -UseBasicParsing
} catch {
    Write-Host "ERROR: Download failed - $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

Write-Host "Extracting..." -ForegroundColor Cyan
try {
    Expand-Archive -Path "$TmpDir\$ZipName" -DestinationPath $TmpDir -Force
} catch {
    Write-Host "ERROR: Extraction failed - $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

$ExtractedExe = Get-ChildItem -Path "$TmpDir\dist" -Filter "git-commit-windows.exe" -Recurse | Select-Object -First 1
if (-not $ExtractedExe) {
    Write-Host "ERROR: Could not find git-commit-windows.exe inside dist/ folder" -ForegroundColor Red
    Get-ChildItem -Path $TmpDir -Recurse | Select-Object FullName
    exit 1
}

Copy-Item $ExtractedExe.FullName -Destination $Target -Force
Write-Host "Installed to $Target" -ForegroundColor Green

if ($env:PATH -notlike "$InstallDir") {
    [Environment]::SetEnvironmentVariable(
        "PATH",
        "$env:PATH;$InstallDir",
        [EnvironmentVariableTarget]::User
    )
    Write-Host "Added $InstallDir to PATH" -ForegroundColor Yellow
}

$env:PATH = [Environment]::GetEnvironmentVariable("PATH","User") + ";" + [Environment]::GetEnvironmentVariable("PATH","Machine")

Remove-Item -Path $TmpDir -Recurse -Force

Write-Host "[OK] $App installed successfully" -ForegroundColor Green