function MoveTo-RecycleBin {
    param (
        [string[]]$Paths
    )

    $shobj = New-Object -ComObject 'Shell.Application'

    foreach ($inputpath in $Paths) {
        $abspath = (Resolve-Path -Path $inputpath).Path
        $parent = Split-Path -Path $abspath -Parent
        $itemname = Split-Path -Path $abspath -Leaf

        Write-Host "Move to Recycle Bin: $inputpath"
        $shobj.Namespace($parent).ParseName($itemname).InvokeVerb('delete')
    }
}