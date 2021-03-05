module fauxDoc

go 1.16

replace FileUtils => ./FileUtils

replace DatabaseUtils => ./DatabaseUtils

require (
	DatabaseUtils v0.0.0-00010101000000-000000000000
	FileUtils v0.0.0-00010101000000-000000000000

)
