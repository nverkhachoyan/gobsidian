---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/4
- monster/environment/forest
- monster/size/large
- monster/type/monstrosity
aliases: ["Girallon"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 152
---
# [Girallon](3-Mechanics\CLI\bestiary\monstrosity/girallon-vgm.md)
*Source: Volo's Guide to Monsters p. 152*  

A girallon looks like an oversized, four-armed ape with gray skin and white fur. Its fangs and claws set it apart from a normal ape, revealing it to be a monstrous predator.

## Forest Hunters

Girallons are most common in temperate or warm forest environments abundant with life. They share the ape's adeptness at climbing, although these half-ton creatures shy away from scaling trees that can't support their bulk. Instead, they stalk the forest floor, lurk in narrow ravines or shallow caves, or hide in ruined sites while waiting for prey to come near. A girallon is surprisingly stealthy, considering its size and its lack of camouflage.

Girallons form loose bands of several individuals and their offspring, usually led by a dominant adult that also tends to be the oldest member of the group. When on the hunt away from their lair, girallons use roars and body language to communicate with one another over distance. Each individual typically hunts alone and widely separated from the others, to ensure that everyone gets adequate fodder. The leader might organize members to work together to make a big kill. If they succeed, everyone in the group shares the spoils, with the best parts going to mothers caring for their young.

In the time since then, numerous creatures have tried to tame, subjugate, or cooperate with the monsters. For instance, yuan-ti enslave girallons, turning them into border sentinels for their serpent kingdoms. Because girallons are known to be peaceful among their own kind, some humanoids have learned how to approach a group's leader, offering food and other gifts in hopes of establishing an alliance with the creatures.

Girallons that are well treated might be willing to serve as guards, though they lack the intelligence to take on tasks more complicated than attacking strangers that enter their domain. If one is taken young and properly trained, a girallon could end up in a seemingly unlikely place, such as guarding the entrance to a city's thieves' guild. Those who would keep a girallon as a pet must always be wary, because the creature could revert to its predatory nature at any time.

## Wall Climbers

The ruins of humanoid habitations, especially those found in deep forests and jungles, seem to attract girallons. They move effortlessly along stairs and balconies, as well as on the sloped rooftops and buttresses of such formations. To a girallon, a city's buildings are just another sort of forest-and better yet, one whose uppermost "branches" can easily support the creatures. In such a setting, the girallons take full advantage of their skill in climbing. The creatures can easily scale walls and battlements, and they perch on tower tops and other high vantages to keep an eye on the surrounding area.

## Magical Origin

The social habits of wild girallons are unusual for apes, as is their instinctive attraction to humanoid structures. These facts, together with the girallon's appearance, lead sages to believe that girallons were created through magic to serve as guardians for some lost empire. When that empire fell ages ago, girallons turned feral and spread out across the world.

```statblock
"name": "Girallon (VGM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "13"
"hp": !!int "59"
"hit_dice": "7d10 + 21"
"stats":
- !!int "18"
- !!int "16"
- !!int "16"
- !!int "5"
- !!int "12"
- !!int "7"
"speed": "40 ft., climb 40 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "3"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": ""
"cr": "4"
"traits":
- "desc": "As a bonus action, the girallon can move up to its speed toward a hostile\
    \ creature that it can see."
  "name": "Aggressive"
- "desc": "The girallon has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
"actions":
- "desc": "The girallon makes five attacks: one with its bite and four with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit. reach 10 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) slashing damage."
  "name": "Claw"
"source":
- "VGM"
- "ToA"
- "IMR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Girallon.webp"
```
^statblock

```encounter-table
name: Girallon
creatures:
 - 1: Girallon
```

## Environment

forest