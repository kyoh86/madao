local collection = {}

local function dump()
    for id, _ in pairs(collection) do
        print(id)
    end
end

return {
    setmetatable(collection, {
        __index = function(original, key)
            if key == "Image" then
                return function(elem)
                    collection[elem.src] = true
                end
            elseif key == "Link" then
                return function(elem)
                    collection[elem.target] = true
                end
            end
        end,
    }),
    {
        Pandoc = function()
            dump()
            return pandoc.Pandoc({}, {})
        end,
    },
}
