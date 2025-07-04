---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/3
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Giff"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 204
---
# [Giff](3-Mechanics\CLI\bestiary\humanoid/giff-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 204*  

## Giff

It's easy to spot the giff in a room: a group of 7-foot-tall, hippopotamus-headed humanoids attired in gaudy military uniforms, with gleaming pistols and muskets on display. These spacefaring mercenaries are renowned for their martial training and their love of explosives.

## Military Organization

Every aspect of giff society is organized along military lines. From birth until death, every giff has a military rank. It must follow orders from those of superior rank, and it can give orders to those of lower rank. Promotions don't depend on age but are granted by a superior as a reward for valor. Giff are devoted to their children, even as most of their education is geared toward fighting and war.

## Mercenaries Extraordinaire

Giff are in high demand as warriors for hire, but they insist on serving in units composed entirely of giff; a giff hiring itself out individually is unheard of. Giff refuse to fight other giff, and will never agree to a contract unless it stipulates that they can sit out a battle rather than wage war against their kin. A giff prizes the reputation of its unit above its own life. Life is fleeting, but the regiment endures for generations or even centuries.

## A Whiff of Gunpowder

Muskets and grenades are the favorite weapons of every giff. The bigger the boom, the brighter the flash, and the thicker the smoke it produces, the more giff love a weapon. Their skill with gunpowder is another reason for their popularity as mercenaries. Giff revel in the challenge of building a bomb big enough to level a fortification. They gladly accept payment in kegs of gunpowder in preference to gold, gems, or other currency.

## No Honor in Magic

Some giff become wizards, clerics, and other kinds of spellcasters, but they're so infrequent that most giff mercenary units have no magical capability. Typical giff are as smart as the average human, but their focus on military training to the exclusion of all other areas of study can make them seem dull-witted to those who have more varied interests.

## Gunpowder by the Keg

Aside from their personal gunpowder weapons, giff ships and mercenary companies carry spare gunpowder in kegs. In an emergency, or any time a large explosion is needed, a whole keg can be detonated. A giff lights the fuse on the keg and can then throw the keg up to 15 feet as part of the same action. The keg explodes at the start of the giff's next turn. Each creature within 20 feet of the exploding keg must make a DC 12 Dexterity saving throw. On a failed save, a creature takes `dice:7d6|text(24)` (`7d6`) fire damage and is knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone). On a successful save, a creature takes half as much damage and isn't knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).

Every other keg of gunpowder within 20 feet of an exploding keg has a 50% chance chance of also exploding. Check each keg only once per turn, no matter how many other kegs explode around it.

```statblock
"name": "Giff (MTF)"
"size": "Medium"
"type": "humanoid"
"alignment": "Lawful Neutral"
"ac": !!int "16"
"ac_class": "[breastplate](/3-Mechanics/CLI/items/breastplate.md)"
"hp": !!int "60"
"hit_dice": "8d8 + 24"
"stats":
- !!int "18"
- !!int "14"
- !!int "17"
- !!int "11"
- !!int "12"
- !!int "12"
"speed": "30 ft."
"senses": "passive Perception 11"
"languages": "Common"
"cr": "3"
"traits":
- "desc": "The giff can try to knock a creature over; if the giff moves at least 20\
    \ feet in a straight line that ends within 5 feet of a Large or smaller creature,\
    \ that creature must succeed on a DC 14 Strength saving throw or take dice:2d6|text(7)\
    \ (2d6) bludgeoning damage and be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Headfirst Charge"
- "desc": "The giff's mastery of its weapons enables it to ignore the loading property\
    \ of muskets and pistols."
  "name": "Firearms Knowledge"
"actions":
- "desc": "The giff makes two pistol attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage, or dice:1d10 + 4|text(9)\
    \ (1d10 + 4) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 40/120 ft.,\
    \ one target. Hit: dice:1d12 + 2|text(8) (1d12 + 2) piercing damage."
  "name": "Musket"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 30/90 ft., one\
    \ target. Hit: dice:1d10 + 2|text(7) (1d10 + 2) piercing damage."
  "name": "Pistol"
- "desc": "The giff throws a grenade up to 60 feet. Each creature within 20 feet of\
    \ the grenade's detonation must make a DC 15 Dexterity saving throw, taking dice:5d6|text(17)\
    \ (5d6) piercing damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Fragmentation Grenade (1/Day)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Giff.webp"
```
^statblock

```encounter-table
name: Giff
creatures:
 - 1: Giff
```

## Environment

urban