---
questObtained: 
questStatus: Not Started
questGiver: "[[Harbin Wester]]"
questLocationObtained: "[[Gnomengarde]]"
questSessionObtained: 
questNotes: 
questLootAvail:
  - "[[clockwork-amulet-xge|Clockwork Amulet]]"
  - "[[pole-of-collapsing-xge|Pole of Collapsing]]"
  - "[[wand-of-pyrotechnics-xge|Wand of Pyrotechnics]]"
  - "[[hat-of-wizardry-xge|Hat of Wizardry]]"
questLootEarned: 
NoteIcon: quest
obsidianUIMode: preview
tags:
  - quest
---

```leaflet
id: gnomengarde_map
lock: true
recenter: true
noScrollZoom: true
image: [[map_gnomengarde.webp]]
### Map Pixel Height x 1 / (Pixels between Bar Scale / 100)
### Map Pixel Width x 1 / (Pixels between Bar Scale / 100) 
### Note that this formula requires adjustments depending on your map. The idea is to determine the number of units between your bar scale. We divide by 100 here because my bar scale measures in 100 units. If your maps scale bar measures in units of 50 them you should divide by 50 instead. The idea is to calculate how many pixels are equal to 1 unit. 
bounds: [[0,0], [204.2222, 320.8889]]
height: 900px
width: 95%
### This sets where the map starts by default. Set it to the middle (half) of your bounds. 
lat: 102.1111
long: 160.4444
### 0 is no zoom. Negative zoom steps away from the map. Positive zoom steps towards the map. 
minZoom: 1.5
### Max zoom is 18. 
maxZoom: 2.5
### Hover mouse over the Reset Zoom icon to see your current zoom level. 
defaultZoom: 2
### How far it zooms in or out with each step. Can be in decimals. 
zoomDelta: 0.5
### This is a string so can be any text. Change it to match your maps measurement scale. 
unit: feet
scale: 1
darkMode: false
```

# `=this.file.name`

> [!infobox]+
> # `=this.file.name`
> ## Quest Details
> Type |  Stat |
> ---|---|
> Date Obtained | `INPUT[datePicker:questObtained]` |
> Status | `INPUT[inlineSelect(option(Not Started), option(In Progress), option(Complete)):questStatus]` |
> Quest Giver | `INPUT[suggester(optionQuery(#npc)):questGiver]` |
> Quest Location | `INPUT[suggester(optionQuery(#Category/Settlement)):questLocationObtained]` |
> Session Obtained | `INPUT[suggester(optionQuery(#journal)):questSessionObtained]` |
> Available Loot | `=this.questLootAvail` |
> Acquired Loot | `=this.questLootEarned` |

> [!readaloud]
> "A clan of reclusive rock gnomes resides in a small network of caves in the mountains to the southeast. The gnomes of Gnomengarde are known for their magical inventions, and they might have something with which to defeat the dragon. Get whatever you can from them. If you bring back something useful and don't want to keep it for yourselves, Townmaster Harbin Wester will pay you 50 gp for it." If the characters undertake this quest, proceed with "Gnomengarde."

### Quest Objective

#### To complete
- [!] Retrieve atleast 1 magic item.

- [?]  Save [[Gnerkeli]] and end [[Korboz]]'s madness.
	- [ ] Retrieve [[clockwork-amulet-xge|Clockwork Amulet]] and [[pole-of-collapsing-xge|Pole of Collapsing]] from grateful gnomes [[Fibblestib]] and [[Dabbledob]].
	- [ ] From [[Gnerkeli]] retrieve [[wand-of-pyrotechnics-xge|Wand of Pyrotechnics]].
	- [ ] From [[Korboz]] retrieve [[hat-of-wizardry-xge|Hat of Wizardry]] .

- [?] Find within the rubble of G13, the Treasury and after successfully searching get both [[clockwork-amulet-xge|Clockwork Amulet]] and [[pole-of-collapsing-xge|Pole of Collapsing]].

### Rewards

-  [[clockwork-amulet-xge|Clockwork Amulet]]
-  [[pole-of-collapsing-xge|Pole of Collapsing]]
-  [[wand-of-pyrotechnics-xge|Wand of Pyrotechnics]]
-  [[hat-of-wizardry-xge|Hat of Wizardry]]

### Additional Information
#### Wild Magic
Any time a spell is cast, a wild magic effect is trigged.

`dice: [](Retrieve%20a%20magical%20weapon!.md#^wild-magic-effect)`

| dice: d20 | Wild Magic Effect                                                                                                                        |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| 1–6       | None                                                                                                                                     |
| 7–10      | The caster's skin turns a vibrant shade of blue.                                                                                         |
| 11–14     | Tiny, insubstantial motes of light circle the caster, shedding bright light in a 10-foot radius and dim light for an additional 10 feet. |
| 15–17     | The caster sprouts wings like those of a butterfly. The wings give the caster a flying speed of 30 feet.                                 |
| 18–19     | The caster teleports up to 60 feet to a random unoccupied space of the DM's choice.                                                      |
| 20        | A whimsical effect of the DM's invention.                                                                                                |
^wild-magic-effect

#### G1. Misty Pool and Mushroom Islands

> [!readaloud]
> In front of you is a river, leading up to  a waterfall. The waterfall and river sits in the middle of two cliff sides with visible entrances into a cave system.  In the middle of the river, stands two small islands which are shrouded in mist from the waterfall.  The islands are covered in massive mushrooms and appear in 3 different colors, red, green and purple.

- [?] The same magic which causes the mushrooms abnormal size is the source of [[wild-magic-tce|Wild Magic]] throughout the [[Gnomengarde]] area.
- [?] Red mushrooms provide oil that the Gnomes use to fuel lanterns.
- [?] Green mushrooms are grounded into flour and used to make a tasty green bread.
- [?] Purple mushrooms are crushed and fermented to make mushroom wine.

#### G2. Waterfall and Rope Bridge

> [!readaloud]
> The waterfall plunges from around 60 feet and obscures a rope bridge, anchored around 20 ft in the air.

- [?] The bridge is "difficult terrain."  Any creature that jumps from it, into the water takes no damage.  The bridge has AC 11, 30 HP and immunity to poison and psychic damage.
- [?] Barrel crab contraptions are too clumsy to cross the bridge without getting tangled in its ropes.

#### G3. Dining Room
> [!readaloud]
> In this room, you find several dining tables and chairs sized for small folk. A stout wooden cabinet against the east wall holds tin dishware and utensils.

#### G4. Kitchen

> [!readaloud]
> Upon entering this room, you witness a small kitchen, furnished with Gnomes in mind.  Everything is either close to the floor or readily reachable by tugging a complicated rope-and-pully mechanism.  You find 5 Gnomes busying themselves.

> [!npc]-  Joybell (Female)
> Before you stands a female Gnome, who is using a poker to stoke the fire of a hot iron stove, standing against the east wall.

>[!npc]- Dimble (Male)
> On the north wall, this male Gnome is using a complicated looking press-like-contraption to squeeze oil out of a big red mushroom and filter the liquid into 4 [[oil-flask|oil flasks]].

>[!npc]- Panana (Female)
>On the west wall, a female Gnome stands atop a low table and uses a mechanical rolling pin contraption to kneed green bread dough.  The severed caps of big green mushrooms are set around her.

>[!npc]- Uppendown (Male)
> A male Gnome is found standing near the center of the room at a table, forming the green dough into loaves.  You can see his tongue sticking out as he carefully shapes each loaf like a master sculptor.

> [!npc]- Tervaround (Female)
> A female Gnome teeters on a stool in the south east corner of the room, stuffing a big purple mushroom into a barrel.
> - [?] Upon being asked, the purple mushrooms are fermented into wine.

- [?] Any character who question the gnomes are urged to speak to [[Fibblestib]] or [[Dabbledob]] in the workshop (G11).

#### G5. Pantry

- [>] This room is piled high with small **wooden creates**.
- [?] Each contains loaves of green mushroom bread and other food related items.

#### G6. Barrel Crabs

- [>] In alcoves, are two gnomish contraptions. 
- [?] Each of them have six-articulated metal legs and a pair of forward-facing pincer claws, they appear to be crab-like.
- [?] A hatch on the top of each barrel opens to reveal an interior compartment equipped with a small, leather-padded seat surrounded by levels, pedals and gears.  The barrels don't appear to be air tight.
- [?] The crablike contraptions serve to move other objects, similar to forklifts.
- [?]  The contraptions are so clumsy that they are useless for delicate work, but they are small enough to navigate [[Gnomengarde]]'s small 5-foot wide passageways.

```statblock
name: Gnomish Barrel Crabs
size: Medium
type: Construct
speed: 15ft
ac: 15
hp: 30
stats:
  - 10
  - 10
  - 10
  - 10
  - 10
  - 10
damage_immunities: "poison, psychic"
actions:
  - name: Claws
    desc: "Melee Weapon Attack: +2 to hit, reach 5 ft., one target. Hit: 2 (1d4) piercing damage and be grappled."
    attack_bonus: 2
    damage_dice: 1d4
```

#### G7. Autoloading Crossbow Platform

> [!scene]+  Facktoré's trouble with the crossbow
> As you enter, you hear screams as crossbow bolts are fired your way.
> "Make it stop, make it stop!!!"
> A contraption, on a rotating platform is spinning out of control.  The contraption is equipped with 4 heavy crossbows, that are firing automatically in every direction.
> Roll `dice: 1d20`   and if the roll is greater than 15, a random PC `dice: 1d3` is hit with a heavy crossbow attack.
> - [?] The party can ignore Facktoré, or they can help stop the contraption by destroying it.
> - [>] Facktoré is a busy female gnomish inventor testing out the crossbow.
> - [?]  She is the inventor of the crossbow contraption and the crab barrel contraptions.
> - [>] She can't string together a coherent sentence to the party after spinning so much.

- [>] Bolted to floor of this room is a **rotating platform**. with 4 heavy crossbows.
- [?] The crossbows are auto reloading.
- [?] A chair mounted 6 feet above the platform, has pedals that causes the entire contraption to spin.  It also has levers that cause the crossbow to fire.

```statblock
name: Gnomish Heavy Crossbow
size: Large 
type: Construct
ac: 13
hp: 45
stats:
  - 10
  - 10
  - 10
  - 10
  - 10
  - 10
damage_immunities: "poison, psychic"
actions:
  - name: Heavy Crossbow
    desc: "Ranged Weapon Attack: +5 to hit, range 50/200ft., one target. Hit: (1d10) piercing damage"
    attack_bonus: 5
    damage_dice: 1d10
reactions:
  - name: When Attacked
    desc: "When it loses 10 hit points, one of its crossbows breaks."
```

#### G8. Mimic and Mushroom Wine

- [>] 12 big barrels are set into wide alcoves.  They are each secured by a wooden brace.  The barrels in the south alcove have been tapped with wooden spigots.
- [?] Two of those barrels are half full, and two are nearly empty.
- [!] Seven of eight barrels in the north and east alcoves are untapped and full of mushroom wine.  Roll `dice: d8|avg|noform` to determine which barrel is a mimic.  When a PC comes nearby, it initiates an encounter.

```encounter
name: Encounter
creatures:
  - "1": Mimic
```

#### G9. Gnome Guard Post

> [!readaloud]+
> Mist from the waterfall dampens this empty cave, which has a 10-foot high ledge overlooking it to the south.

- [?] The ledge can be reached by following the curved tunnel to the east or by scaling the slick rock with a successful DC 12 Athletics (Strength) check.
- [>] Two Gnomes stand on the ledge.

> [!npc]+ Ulla (Male)
> When anyone enters, Ulla calls out "Who goes there?" in Gnomish.

> [!npc]+ Pog (Female)
> When anyone enters, Pog repeats "Who goes there?" but in Common.

- [?] Both Pog and Ulla have orders to attack "shapeshifters on sight."
- [?] PCs can prove they are not shapeshifters with some checks. 
	- [ ] DC 10 Charisma [Deception](/3-Mechanics/CLI/rules/skills.md#Deception)
	- [ ] DC 10 Charisma [Intimidation](/3-Mechanics/CLI/rules/skills.md#Intimidation) 
	- [ ] DC 10 Charisma [Persuasion](/3-Mechanics/CLI/rules/skills.md#Persuasion)

#### G10. Spinning Blades

- [>] This room appears to be obscured by the mist from the waterfall.  You can hear the mechanical rotation of some kind of contraption nearby.
- [?]  The larger eastern part of the room contains two rapidly spinning devices that look like turnstiles fitted with stacks of long, sharp blades spaced 1 foot apart.
- [!] The northern turnstile spins counterclockwise, while the southern one spins clockwise. 
	- [?]  On entry to this area, creatures must make a DC 15 Dexterity saving throw.  Take `dice: 4d8` slashing damage on a failed save, or half as much on a successful one.	
- [?] On the south wall of the western part of the room is a lever that stops the spinning of the turnstiles.

#### G11. Inventors' Workshop

- [>] As characters approach, they hear an argument between two Gnomes.

>[!npc]+ Fibblestib (Male)
> Fibblestib is found arguing to **Dabbledob** about an invention that will cure King Korboz's madness.
> Fibblestib proposes "sanity ray" as the cure.

>[!npc]+ Dabbledob (Female)
>While arguing, Dabbledob proposes to build something called a "straitjacket."

- [>] On approach, the two Gnomes fill in the players about what has been happening to the Korboz and to rescue King Gnerkli, they also talk about a shapeshifter amongst the Gnomes as a threat, but it is secondary.
- [>] The two Gnomes also talk about magic items as rewards for the PC's help.

> [!readaloud]
> The workshop is cluttered with half-completed gnomish inventions that serve no purpose, as well as worktables strewn with tinker's tools.

- [>] A 10-foot-high ledge overlooks the room, set with a pedestal.
- [?] On the pedestal is a spellbook that the Gnomes share.  It's title is *Magick of Gnomeengarde* and it contains the wizard spells [[burning-hands|Burning Hands]], [[detect-magic|Detect Magic]], [[identify|Identify]], [[mage-armor|Mage Armor]], [[magic-missile|Magic Missile]], [[3-Mechanics/CLI/spells/shield|Shield]] and [[sleep|Sleep]].

#### G12. Gnome Domiciles

- [>] The floor of this cave is strewn with the remnants of old campfires.  Four side caves serve as sleeping areas, with five small wooden cots crammed into each one.
- [>] **Eight** Gnomes sleep soundly here, with two Gnomes in each side cave.

> [!npc]+ Caramip (Female)

> [!npc]+ Jabby (Female)

> [!npc]+ Nyx (Female)

> [!npc]+ Quibby (Female)

> [!npc]+ Anverth  (Male)

> [!npc]+ Delebean (Male)

> [!npc]+ Pallabar (Male)

> [!npc]+ Zook (Male)

- [?] Characters can move through this area without waking the Gnomes up.  If woken, they refer the PCs to Dabbledob and Fiddlestib to talk about their troubles.

#### G13. Treasury

- [?] The door to this room is locked.   Fibblestib and Dabbledob carry the keys.
- [?] Amid the clutter of metals in the room, PCs can find a [[clockwork-amulet-xge|Clockwork Amulet]] and a [[pole-of-collapsing-xge|Pole of Collapsing]].  They can be found with [[detect-magic|Detect Magic]], but otherwise it takes 1 hour to find each item.

#### G14. Throne Room

- [>] Situated atop a stone dais are two squat thrones made of scrap metal and sized for gnomes.
- [?] A secret door in the north wall conceals a short tunnel leading to area G15. Only the gnome kings know of this secret passage.


#### G15. Gnome Kings' Bedroom

- [?] King Korboz has locked himself and King Gnerkli in their bedroom.
	- [?] There is a secret door that others could find and use to gain entry.
- [?] If PCs knocke on the door, or announce arrival, Korboz speaks to them from inside and warns of a "shapechanger" in their midst.
- [?] Korboz doesn't regain his senses until the PCs assure him that the monster has been found and killed.
	- [?] DC 12 (Persuasion) Charisma check to convince him, whether or not its actually dead.

- [>] Korboz and Gnerkli are Gnomes, each wearing a jagged metal crown and a patchwork cloack. 
- [>] Gnerkli is glued to a chair and restrained.
- [?] Korboz carries a flask of solvent that dissolves the glue on contact. 

- [?] A small unlocked chest under the gnomes' bed contains a [[hat-of-wizardry-xge|Hat of Wizardry]] and a fully charged [[wand-of-pyrotechnics-xge|Wand of Pyrotechnics]].