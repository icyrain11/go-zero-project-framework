@ECHO OFF

SET IDL_PATH=server/api/idl/
SET GEN_PATH=server/api/
SET SWAGGER_PATH=server/api/idl/

goctl api go -api %IDL_PATH%api.api -dir %GEN_PATH%
goctl api plugin -plugin goctl-swagger="swagger -filename api.json" -api %IDL_PATH%api.api -dir %SWAGGER_PATH%

ECHO Api CodeGen Done.