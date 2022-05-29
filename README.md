

## Pokemon Go!!

### Terminal-based Pokemon game, created in golang. Not the AR game. 




## TODO
- [x] Set up global pointers
- [x] Make an output window for console logging
- [ ] more extensive logging system (write to file)
- [ ] Message queue for console logging

- [x] Refurnish the map system
- [ ] Going into different loading zones
- [ ] Grass and RNG pokemon spawning
- [ ] Battle mode
- [ ] NPCs, basically a single-pixel building
- [ ] Rendering sprite frames


## Bugs
- [x] collision is messed up (urgent fix)
- [ ] console text gets shifted when the player is held in place by obstacles


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



