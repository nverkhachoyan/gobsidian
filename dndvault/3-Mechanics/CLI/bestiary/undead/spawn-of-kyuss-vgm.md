---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/desert
- monster/environment/underdark
- monster/size/medium
- monster/type/undead
aliases: ["Spawn of Kyuss"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 192
---
# [Spawn of Kyuss](3-Mechanics\CLI\bestiary\undead/spawn-of-kyuss-vgm.md)
*Source: Volo's Guide to Monsters p. 192*  

Kyuss was a high priest of Orcus who plundered corpses from necropolises to create the first spawn of Kyuss. Even centuries after Kyuss's death, his mad disciples continue performing the horrific rites he perfected.

## Plague of Worms

From a distance or in poor light, a spawn of Kyuss looks like an ordinary zombie. As it comes into clearer view, one can see scores of little green worms crawling in and out of it. These worms jump onto nearby humanoids and burrow into their flesh. A worm that penetrates a humanoid body makes its way to the creature's brain. Once inside the brain, the worm kills its host and animates the corpse, transforming it into a spawn of Kyuss that breeds more worms. The dead humanoid's soul remains trapped inside the corpse, preventing the individual from being raised or resurrected until the undead body is destroyed. The horror of being a soul imprisoned in an undead body drives a spawn of Kyuss insane.

## Corruption Without End

Spawn of Kyuss are expressions of Orcus's intent to replace all life with undeath. Left to its own devices, a solitary spawn of Kyuss travels aimlessly. If it stumbles across a living creature, the spawn attacks with the sole intent of creating more spawn. Whether they are dispersed or clustered, spawn reproduce exponentially if nothing stops them.

## Undead Nature

Spawn of Kyuss require no air, food, drink, or sleep.

```statblock
"name": "Spawn of Kyuss (VGM)"
"size": "Medium"
"type": "undead"
"alignment": "Chaotic Evil"
"ac": !!int "10"
"hp": !!int "76"
"hit_dice": "9d8 + 36"
"stats":
- !!int "16"
- !!int "11"
- !!int "18"
- !!int "5"
- !!int "7"
- !!int "3"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "1"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 8"
"languages": "understands the languages it knew in life but can't speak"
"cr": "5"
"traits":
- "desc": "The spawn of Kyuss regains 10 hit points at the start of its turn if it\
    \ has at least 1 hit point and isn't in sunlight or a body of running water. If\
    \ the spawn takes acid, fire, or radiant damage, this trait doesn't function at\
    \ the start of the spawn's next turn. The spawn is destroyed only if it starts\
    \ its turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
- "desc": "If the spawn of Kyuss is targeted by an effect that cures disease or removes\
    \ a curse, all the worms infesting it wither away, and it loses its Burrowing\
    \ Worm action."
  "name": "Worms"
"actions":
- "desc": "The spawn of Kyuss makes two attacks with its claws and uses Burrowing\
    \ Worm."
  "name": "Multiattack"
- "desc": "A worm launches from the spawn of Kyuss at one humanoid that the spawn\
    \ can see within 10 feet of it. The worm latches onto the target's skin unless\
    \ the target succeeds on a DC 11 Dexterity saving throw. The worm is a Tiny undead\
    \ with AC 6, 1 hit point, a 2 (-4) in every ability score, and a speed of 1 foot.\
    \ While on the target's skin, the worm can be killed by normal means or scraped\
    \ off using an action (the spawn can use this action to launch a scraped-off worm\
    \ at a humanoid it can see within 10 feet of the worm). Otherwise, the worm burrows\
    \ under the target's skin at the end of the target's next turn, dealing 1 piercing\
    \ damage to it. At the end of each of its turns thereafter, the target takes dice:2d6|text(7)\
    \ (2d6) necrotic damage per worm infesting it (maximum of dice: 10d6|avg|noform\
    \ (10d6)). A worm-infested target dies if it drops to 0 hit points, then rises\
    \ 10 minutes later as a spawn of Kyuss. If a worm-infested creature is targeted\
    \ by an effect that cures disease or removes a curse, all the worms infesting\
    \ it wither away."
  "name": "Burrowing Worm"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) slashing damage plus dice:2d6|text(7)\
    \ (2d6) necrotic damage."
  "name": "Claw"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Spawn%20of%20Kyuss.webp"
```
^statblock

```encounter-table
name: Spawn of Kyuss
creatures:
 - 1: Spawn of Kyuss
```

## Environment

underdark, desert