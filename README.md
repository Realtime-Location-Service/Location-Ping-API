# Location-REST-API
When I start programming back in 2011-ish, I start to hear DHL started showing realtime location of a parcel. From that time, now a days, you see a lot of applications (specially, rider sharing, home delivery, kids/student tracking) are geared with realtime location tracking. Hence, its became very common in our ecosystem. That's good but a bit bad from developer point of view. Development team of all those applications are solving same common problem duplicatedly like:
- saving and serving location
- how to scale when it starts getting 10K req/sec
- how to optimize storage when location data is exceeding Giga to Tera
- how to extrapolate a location when it is off-road or missing for some long minutes
- how to calculate arrival time more efficiently
- how to increase location accuracy in case of absense / low-performance of GPS
- etc

Why not solve all those problem in a single place? In that way, development team can focus on only the mission-critical things or more other creative stuffs :) This repository set is aimed for that purpose (e.g., solve in 1 place). Hope you help me.

# How to run
`./run.sh`

# How to test
`go test ./... -v`

