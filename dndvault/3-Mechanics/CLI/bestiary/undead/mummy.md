---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/3
- monster/environment/desert
- monster/size/medium
- monster/type/undead
aliases: ["Mummy"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 228, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.
---
# [Mummy](3-Mechanics\CLI\bestiary\undead/mummy.md)
*Source: Monster Manual p. 228, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.*  

Raised by dark funerary rituals, a mummy shambles from the shrouded stillness of a time-lost temple or tomb. Having been awoken from its rest, it punishes transgressors with the power of its unholy curse.

## Preserved Wrath

The long burial rituals that accompany a mummy's entombment help protect its body from rot. In the embalming process, the newly dead creature's organs are removed and placed in special jars, and its corpse is treated with preserving oils, herbs, and wrappings. After the body has been prepared, the corpse is typically wrapped in linen bandages.

## The Will of Dark Gods

An undead mummy is created when the priest of a death god or other dark deity ritually imbues a prepared corpse with necromantic magic. The mummy's linen wrappings are inscribed with necromantic markings before the burial ritual concludes with an invocation to darkness. As a mummy endures in undeath, it animates in response to conditions specified by the ritual. Most commonly, a transgression against its tomb, treasures, lands, or former loved ones will cause a mummy to rise.

## The Punished

Once deceased, an individual has no say in whether or not its body is made into a mummy. Some mummies were powerful individuals who displeased a high priest or pharaoh, or who committed crimes of treason, adultery, or murder. As punishment, they were cursed with eternal undeath, embalmed, mummified, and sealed away. Other times, mummies acting as tomb guardians are created from slaves put to death specifically to serve a greater purpose.

## Creature of Ritual

A mummy obeys the conditions and parameters laid down by the rituals that created it, driven only to punish transgressors. The overwhelming terror that foreshadows a mummy's attack can leave the intended victim paralyzed with fright. In the days following a mummy's touch, a victim's body rots from the outside in, until nothing but dust remains.

## Ending a Mummy's Curse

Rare magic can undo or dispel the ritual that gave rise to a mummy, allowing it to truly die. More commonly, a mummy can be sent back to its endless rest by undoing the transgression that caused it to rise. A sacred idol might be replaced in its niche, a stolen treasure could be returned to its tomb, or a temple might be purified of despoiling bloodshed.

More ephemeral or permanent offenses, such as revealing a secret the mummy wished kept or killing an individual the mummy loved, can't be so easily remedied. In such cases, a mummy might slaughter all the creatures responsible and still not sate its wrath.

## Undead Archives

Though they seldom bother to do so, mummies can speak. As a result, some serve as undead repositories of lost lore, and can be consulted by the descendants of those who created them. Powerful individuals sometimes intentionally sequester mummies away for occasional consultation.

## Undead Nature

A mummy doesn't require air, food, drink, or sleep.

> [!quote]- A quote from X the Mystic's 7th rule of dungeon survival  
> 
> Before opening a sarcophagus, light a torch.


```statblock
"name": "Mummy"
"size": "Medium"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "11"
"ac_class": "natural armor"
"hp": !!int "58"
"hit_dice": "9d8 + 18"
"stats":
- !!int "16"
- !!int "8"
- !!int "15"
- !!int "6"
- !!int "10"
- !!int "12"
"speed": "20 ft."
"saves":
  "Wisdom": !!int "2"
"damage_vulnerabilities": "fire"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "the languages it knew in life"
"cr": "3"
"actions":
- "desc": "The mummy can use its Dreadful Glare and makes one attack with its rotting\
    \ fist."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) bludgeoning damage plus dice:3d6|text(10)\
    \ (3d6) necrotic damage. If the target is a creature, it must succeed on a DC\
    \ 12 Constitution saving throw or be cursed with mummy rot. The cursed target\
    \ can't regain hit points, and its hit point maximum decreases by dice:3d6|text(10)\
    \ (3d6) for every 24 hours that elapse. If the curse reduces the target's hit\
    \ point maximum to 0, the target dies, and its body turns to dust. The curse lasts\
    \ until removed by the [remove curse](/3-Mechanics/CLI/spells/remove-curse.md)\
    \ spell or other magic."
  "name": "Rotting Fist"
- "desc": "The mummy targets one creature it can see within 60 feet of it. If the\
    \ target can see the mummy, it must succeed on a DC 11 Wisdom saving throw against\
    \ this magic or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ until the end of the mummy's next turn. If the target fails the saving throw\
    \ by 5 or more, it is also [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for the same duration. A target that succeeds on the saving throw is immune\
    \ to the Dreadful Glare of all mummies (but not mummy lords) for the next 24 hours."
  "name": "Dreadful Glare"
"source":
- "MM"
- "RoT"
- "ToA"
- "WDH"
- "WDMM"
- "GoS"
- "BGDIA"
- "IDRotF"
- "TCE"
- "CM"
- "DSotDQ"
- "PSI"
- "PaBTSO"
- "AATM"
- "BMT"
- "DoDk"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Mummy.webp"
```
^statblock

```encounter-table
name: Mummy
creatures:
 - 1: Mummy
```

## Environment

desert