mingw使用msvc的dll

用pexports把dll中的符号导出来

pexports libmyadd.dll > libmyadd.def

然后用mingw中的dlltool把dl转成.a

dlltool.exe -D .\Mv3dLp.dll -d .\Mv3dLp.def -l libMv3dLp.dll.a -k
