---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/forest
- monster/size/small
- monster/type/humanoid/grung
aliases: ["Grung Elite Warrior"]
NoteIcon: monster
BestiaryType: humanoid (grung)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 157
---
# [Grung Elite Warrior](3-Mechanics\CLI\bestiary\humanoid/grung-elite-warrior-vgm.md)
*Source: Volo's Guide to Monsters p. 157*  

Grungs are aggressive frog-like humanoids found in rain forests and tropical jungles. They are fiercely territorial and see themselves as superior to most other creatures.

## Tree-Dwelling Amphibians

Grungs live in trees and prefer shade. A grung hatchery is maintained in well-guarded ground-level pools. About three months after hatching, a grung tadpole takes on the shape of an adult. It takes another six to nine months for a grung juvenile to reach maturity.

Green grungs are the tribe's warriors, hunters, and laborers, and blue grungs work as artisans and in other domestic roles. Supervising and guiding both groups are the purple grungs, which serve as administrators and commanders. (Use the grung stat block to represent members of the green, blue, and purple castes.)

Red grungs are the tribe's scholars and magic users. They are superior to purple, blue, and green grungs and given proper respect even by grungs of higher status. (Use the grung wildling stat block to represent members of the red caste.)

Higher castes include orange grungs, which are elite warriors that have authority over all lesser grungs, and gold grungs, which hold the highest leadership positions. A tribe's sovereign is always a gold grung. (Use the grung elite warrior stat block to represent members of the orange and gold castes.)

A grung normally remains in its caste for life. On rare occasions, an individual that distinguishes itself with great deeds can earn an invitation to join a higher caste. Through a combination of herbal tonics and ritual magic, an elevated grung changes color and is inducted into its new caste in the same way that a juvenile of the caste would be. From then on, the grung and its progeny are members of the higher caste.

Sentient, poisonous frogs that live in trees. Truly, the gods hate us.

-Volo

## Castes and Colors

Grung society is a caste system. Each caste lays eggs in a separate hatching pool, and juvenile grungs join their caste upon emergence from the hatchery. All grungs are a dull greenish gray when they are born, but each individual takes on the color of its caste as it grows to adulthood.

## Naturally Toxic

All grungs secrete a substance that is harmless to them but poisonous to other creatures. A grung also uses venom to poison its weapons.

## Slavers

Grungs are always on the lookout for creatures they can capture and enslave. Grungs use slaves for all manner of menial tasks, but mostly they just like bossing them around. Slaves are fed mildly [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) food to keep them lethargic and compliant. A creature afflicted in this way over a long period of time becomes a shell of its former self and can be restored to normalcy only by magic.

## Water Dependency

A grung that fails to immerse itself in water for at least 1 hour during a day suffers one level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion) at the end of that day. A grung can recover from this [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion) only through magic or by immersing itself in water for at least 1 hour.

## Variant: Grung Poison

Grung poison loses its potency 1 minute after being removed from a grung. A similar breakdown occurs if the grung dies.

A creature [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) by a grung can suffer an additional effect that varies depending on the grung's skin color. This effect lasts until the creature is no longer [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) by the grung.

### Green

The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) creature can't move except to climb or make standing jumps. If the creature is flying, it can't take any actions or reactions unless it lands.

### Blue

The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) creature must shout loudly or otherwise make a loud noise at the start and end of its turn.

### Purple

The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) creature feels a desperate need to soak itself in liquid or mud. It can't take actions or move except to do so or to reach a body of liquid or mud.

### Red

The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) creature must use its action to eat if food is within reach.

### Orange

The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) creature is [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) of its allies.

### Gold

The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) creature is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) and can speak Grung.

```statblock
"name": "Grung Elite Warrior (VGM)"
"size": "Small"
"type": "humanoid"
"subtype": "grung"
"alignment": "Lawful Evil"
"ac": !!int "13"
"hp": !!int "49"
"hit_dice": "9d6 + 18"
"stats":
- !!int "7"
- !!int "16"
- !!int "15"
- !!int "10"
- !!int "11"
- !!int "12"
"speed": "25 ft., climb 25 ft."
"saves":
  "Dexterity": !!int "5"
"skillsaves":
  "Athletics": !!int "2"
  "Stealth": !!int "5"
  "Perception": !!int "2"
  "Survival": !!int "2"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "passive Perception 12"
"languages": "Grung"
"cr": "2"
"traits":
- "desc": "The grung can breathe air and water."
  "name": "Amphibious"
- "desc": "Any creature that grapples the grung or otherwise comes into direct contact\
    \ with the grung's skin must succeed on a DC 12 Constitution saving throw or become\
    \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for 1 minute. A [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ creature no longer in direct contact with the grung can repeat the saving throw\
    \ at the end of each of its turns, ending the effect on itself on a success."
  "name": "Poisonous Skin"
- "desc": "The grung's long jump is up to 25 feet and its high jump is up to 15 feet,\
    \ with or without a running start."
  "name": "Standing Leap"
"actions":
- "desc": "Melee or Ranged Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing\
    \ damage, and the target must succeed on a DC 12 Constitution saving throw or\
    \ take dice:2d4|text(5) (2d4) poison damage."
  "name": "Dagger"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage, and the\
    \ target must succeed on a DC 12 Constitution saving throw or take dice:2d4|text(5)\
    \ (2d4) poison damage."
  "name": "Shortbow"
- "desc": "The grung makes a chirring noise to which grungs are immune. Each humanoid\
    \ or beast that is within 15 feet of the grung and able to hear it must succeed\
    \ on a DC 12 Wisdom saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the end of the grung's next turn."
  "name": "Mesmerizing Chirr (Recharge 6)"
"source":
- "VGM"
- "ToA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Grung%20Elite%20Warrior.webp"
```
^statblock

```encounter-table
name: Grung Elite Warrior
creatures:
 - 1: Grung Elite Warrior
```

## Environment

forest