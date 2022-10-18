# Collisions

## Ship to gound collisions

Originally I thought to define the 'caves' in the game as single, large concave polygon and use the ray-casting method for each vertex of the ship to determine whether the ship was 'in space' or 'in the wall'.

The single large polygon approach doesn't sit well with me now however - to make interesting and complex cave geometries the polygon will have to have a lot of vertices. The more complex the polygon becomes, the more vertices we have to test for the ray casting solution.

Instead, If I adopt a tile-based approach, with each cave tile composing a simple, convex polygon and only testing tiles that are adjacent to the player position, then I believe it will be simpler and faster to test for collisions, as well as allowing the map to be any size.

Cave tiles



