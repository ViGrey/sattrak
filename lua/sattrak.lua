local socket = require("socket.core")

local injectionContent = ""

local i = 0;
local controller1 = 0x501;
local next = next;

local HOST, PORT = "localhost", 25544

local sock = nil

function handleConnRead()
  if sock ~= nil then
    local _, err, part = sock:receive("*all")
    if err == "closed" then
      createSock()
    elseif part and #part > 0 then
      injectionContent = injectionContent .. part
    end
  else
    createSock()
  end
end

function connect(address, port, laddress, lport)
  local sock, err = socket.tcp()
  if not sock then
    return nil, err
  end
  if laddress then
    local res, err = sock:bind(laddress, lport, -1)
    if not res then
      return nil, err
    end
  end
  local res, err = sock:connect(address, port)
  if not res then
    return nil, err
  end
  return sock, nil
end

function createSock()
  sock, err = connect(HOST, PORT)
  if sock ~= nil then
    sock:settimeout(0)
  end
end

function handleControllerRead()
  if injectionContent == "" then
    memory.writebyte(controller1, 0);
  else
    local injectionByte = string.byte(string.sub(injectionContent, 1, 1))
    memory.writebyte(controller1, injectionByte)
    injectionContent = string.sub(injectionContent, 2);
  end
end

function main()
  while true do
    handleConnRead()
    emu.frameadvance()
  end
end

memory.registerexec(0x801B, handleControllerRead);
createSock()

main();
