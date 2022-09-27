-- sets up a computercraft turtle to become an atm -- 
resp = http.get("https://raw.githubusercontent.com/rxi/json.lua/master/json.lua")
jsonLibText = resp.readAll()
jsonLibFile = fs.open("json.lua", "w")
jsonLibFile.writeString(jsonLibFile)
jsonLibFile.close()