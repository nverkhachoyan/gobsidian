---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vrgr
- monster/cr/3
- monster/size/tiny
- monster/type/monstrosity
aliases: ["Carrion Stalker"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Van Richten's Guide to Ravenloft p. 230
---
# [Carrion Stalker](3-Mechanics\CLI\bestiary\monstrosity/carrion-stalker-vrgr.md)
*Source: Van Richten's Guide to Ravenloft p. 230*  

A carrion stalker begins life as a pale larva that infests a corpse. Over the course of weeks, this grub burrows, feeds, and grows, ultimately developing into a chitinous mass of pincers and tentacles. When an adult carrion stalker detects movement, it bursts from its corpse-cradle to attack, intent on implanting its young into the living and starting its species' life cycle anew.

More than one necromancer has animated a corpse infested with carrion stalker larvae. While this can prove shocking and deadly, some depraved spellcasters cultivate carrion stalkers within zombies. The embedded carrion stalkers ride along in their freshly animated conveyances, bursting forth once they detect living creatures nearby. This destroys the zombie, but unleashes a new horror.

Carrion stalkers also enjoy symbiotic relationships with carrion crawlers. Carrion crawlers won't devour bodies infested by carrion stalkers, but they often pick up stalker larvae as they root among filth. The crawlers then spread these grubs, potentially infecting whole sewers, graveyards, or battlefields with carrion stalkers. In return, carrion stalkers avoid preying on carrion crawlers.

```statblock
"name": "Carrion Stalker (VRGR)"
"size": "Tiny"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "35"
"hit_dice": "10d4 + 10"
"stats":
- !!int "6"
- !!int "16"
- !!int "12"
- !!int "2"
- !!int "13"
- !!int "6"
"speed": "30 ft., burrow 30 ft."
"skillsaves":
  "Stealth": !!int "7"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)"
"senses": "tremorsense 60 ft., passive Perception 11"
"languages": ""
"cr": "3"
"actions":
- "desc": "The carrion stalker makes three Tentacle attacks. If it is attached to\
    \ a creature, it can replace one Tentacle attack with Larval Burst, if available."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage, and the carrion stalker\
    \ attaches to the target and pulls itself into the target's space. While attached,\
    \ the carrion stalker moves with the target and has advantage on attack rolls\
    \ against it.\n\nA creature can use its action to try to detach the carrion stalker\
    \ and force it to move into the nearest unoccupied space, doing so with a successful\
    \ DC 11 Strength check. On its turn, the carrion stalker can detach itself from\
    \ the target by using 5 feet of movement. When it dies, the carrion stalker detaches\
    \ from any creature it is attached to."
  "name": "Tentacle"
- "desc": "The carrion stalker releases a burst of larvae in a 10-foot-radius sphere\
    \ centered on itself. Each creature in that area must succeed on a DC 13 Constitution\
    \ saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned).\
    \ A creature [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this\
    \ way takes dice:2d6|text(7) (2d6) poison damage at the start of each of its\
    \ turns as larvae infest its body. The creature can repeat the saving throw at\
    \ the end of each of its turns, ending the effect on itself on a success. Any\
    \ effect that cures disease or removes the [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ condition instantly kills the larvae in the creature, ending the effect on it.\n\
    \nIf a creature is reduced to 0 hit points by the infestation, it dies. The larvae\
    \ remain in the corpse, and one survives to become a fully grown carrion stalker\
    \ in dice: 1d4|avg|noform (1d4) weeks. Any effect that cures diseases or removes\
    \ the [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) condition that\
    \ targets the corpse instantly kills the larvae."
  "name": "Larval Burst (1/Day)"
"source":
- "VRGR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VRGR/Carrion%20Stalker.webp"
```
^statblock

```encounter-table
name: Carrion Stalker
creatures:
 - 1: Carrion Stalker
```