local collection = {}

local function collect(elem)
    if not elem.identifier then
        return
    end
    collection[elem.identifier] = true
end

local function dump()
    for id, _ in pairs(collection) do
        print(id)
    end
end

return {
    setmetatable(collection, {
        __index = function()
            return collect
        end,
    }),
    {
        Pandoc = function()
            dump()
            return pandoc.Pandoc({}, {})
        end,
    },
}
