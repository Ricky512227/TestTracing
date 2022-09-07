Dev will generate lib file and copy to any path in the node.

Syntax ::  applyPatch.xx "absolutepath of the patch" "servicename"

1. Get the md5sum of the patchFile.
2. Check the patch file is already present in the /tcnVol, 
   1.If not, 
        1. Copy the patch to the /tcnVol of the required service.
        2. If the library is already exists in the /opt/SMAW/INTP/lib64 and md5sum of the patchfile is diff.
            a. Take the batkup of the existing lib file, link the new file form /tcnVol to the /opt/SMAW/INTP/lib64.
  2. If yes,
        1. the md5sum are diff and the same patch is already linked form /tcnVol to the /opt/SMAW/INTP/lib64.
        2. then Copy the patch to the /tcnVol of the required service
3. Login the required service of mcc container, kill the service process
4. Monitor till the process comes up.
5. If all the process comesup then, 
    1. Log Patching as successfull
    2. else, Patching as unsuccesfull.
    
 
      
  

