

## Pokemon Go!!

### Terminal-based Pokemon game, created in golang. Not the AR game. 


## Run instructions
- Run `go run main.go` in the project root


## TODO
- [x] Set up global pointers
- [x] Make an output window for console logging
- [x] more extensive logging system (write to file)
- [x] Message queue for console logging

- [x] Refurnish the map system
- [x] Going into different loading zones
- [ ] Store current room state
- [ ] Make players unable to step out of rooms unless by portal
- [ ] Grass and RNG pokemon spawning
- [ ] Battle mode
- [ ] NPCs, basically a single-pixel building
- [ ] Rendering sprite frames


## Bugs
- [x] collision is messed up (urgent fix)
- [x] console text gets shifted when the player is held in place by obstacles


Entity
- dialogue
- trigger battle

  NPC
  - heal
  - give item
  - buy/sell items

  Pokemon
  - wrapper for special text before and after battle

Battle


### Items
- item id, quantity
- display message


### NPC
  - contains a list of actions
  - has already talked to?
  - walk pattern

#### NPC actions
- give item
- trigger battle
- dialogue
- heal
- buy/sell items


### Terrain/tiles
- grass (spawn pokemon)
- portal (doors, stairs)
- obstacles
- pickup items
- NPCs


### Inventory



