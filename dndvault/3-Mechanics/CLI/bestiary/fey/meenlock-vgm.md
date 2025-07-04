---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/forest
- monster/environment/swamp
- monster/environment/urban
- monster/size/small
- monster/type/fey
aliases: ["Meenlock"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 170
---
# [Meenlock](3-Mechanics\CLI\bestiary\fey/meenlock-vgm.md)
*Source: Volo's Guide to Monsters p. 170*  

Meenlocks are deformed fey that invoke terror and seek to destroy all that is good, innocent, and beautiful. They primarily live in forests, although they adapt well to urban and subterranean settings.

## Fear Incarnate

Meenlocks are spawned by fear. Whenever fear overwhelms a creature in the Feywild, or in any other location where the Feywild's influence is strong, one or more meenlocks might spontaneously arise in the shadows or darkness nearby. If more than one meenlock is born, a lair also magically forms. The earth creaks and moans as narrow, twisting tunnels open up within it. One of these newly formed passageways serves as the lair's only entrance and exit.

Meenlocks give other creatures the creeps and project a supernatural aura that instills terror in those nearby. So evil and twisted are they that a palpable sense of foreboding haunts those who intrude upon a meenlock lair. Inside the warren, black moss covers every surface, muffling sound. A large central chamber serves as the meenlocks' den, where they torment captives.

During the day, meenlocks confine themselves to their dark warrens. At night, they crawl out of their tunnels to torment sleeping prey, particularly those who seem to embody all that is good in the world. Meenlocks like to paralyze creatures with their claws, drag them back to their hidden den, beat them unconscious, and telepathically torture them over a period of hours. A humanoid that succumbs to this psychic torment undergoes a transformation into an evil, full-grown meenlock (see the "Telepathic Torment" sidebar).

## Dark Dwellers

A meenlock shuns bright light. It can supernaturally sense areas of darkness and shadow in its vicinity and thus is able to teleport from one darkened space to another-enabling it to sneak up on its prey or run away when outmatched.

## Telepathic Tormentors

Meenlocks have no form of communication other than telepathy. They can use it to project unsettling hallucinations into the minds of their prey. These hallucinations take the form of terrible whispers or fleeting movements just at the edges of one's peripheral vision.

> [!note] Telepathic Torment
> 
> Up to four meenlocks can telepathically torment one [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated) creature, filling its mind with disturbing sounds and dreadful imagery. Participating meenlocks can't use their telepathy for any other purpose during this time, though they can move about and take actions and reactions as normal. This torment has no effect on a creature that is immune to the [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) condition. If the creature is susceptible and remains [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated) for 1 hour, the creature must make a Wisdom saving throw, taking `dice:3d6|text(10)` (`3d6`) psychic damage on a failed save, or half as much damage on a successful one. The save DC is 10 + the number of meenlocks participating in the torment, considering only those that remain within sight of the victim for the entire hour and aren't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated) during it. The process can be repeated. A humanoid that drops to 0 hit points as a result of this damage instantly transforms into a meenlock at full health and under the DM's control. Only a [wish](/3-Mechanics/CLI/spells/wish.md) spell or divine intervention can restore a transformed creature to its former state.
^telepathic-torment

```statblock
"name": "Meenlock (VGM)"
"size": "Small"
"type": "fey"
"alignment": "Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "31"
"hit_dice": "7d6 + 7"
"stats":
- !!int "7"
- !!int "15"
- !!int "12"
- !!int "11"
- !!int "10"
- !!int "8"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "4"
  "Survival": !!int "2"
"condition_immunities": "[frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 120 ft., passive Perception 14"
"languages": "telepathy 120 ft."
"cr": "2"
"traits":
- "desc": "Any beast or humanoid that starts its turn within 10 feet of the meenlock\
    \ must succeed on a DC 11 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ until the start of the creature's next turn."
  "name": "Fear Aura"
- "desc": "While in bright light, the meenlock has disadvantage on attack rolls, as\
    \ well as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight."
  "name": "Light Sensitivity"
- "desc": "As a bonus action, the meenlock can teleport to an unoccupied space within\
    \ 30 feet of it, provided that both the space it's teleporting from and its destination\
    \ are in dim light or darkness. The destination need not be within line of sight."
  "name": "Shadow Teleport (Recharge 5-6)"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 2|text(7) (2d4 + 2) slashing damage, and the target must\
    \ succeed on a DC 11 Constitution saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. The target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success."
  "name": "Claws"
"source":
- "VGM"
- "CM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Meenlock.webp"
```
^statblock

```encounter-table
name: Meenlock
creatures:
 - 1: Meenlock
```

## Environment

forest, swamp, urban