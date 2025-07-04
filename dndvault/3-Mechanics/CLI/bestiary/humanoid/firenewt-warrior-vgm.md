---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-2
- monster/environment/desert
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/firenewt
aliases: ["Firenewt Warrior"]
NoteIcon: monster
BestiaryType: humanoid (firenewt)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 142
---
# [Firenewt Warrior](3-Mechanics\CLI\bestiary\humanoid/firenewt-warrior-vgm.md)
*Source: Volo's Guide to Monsters p. 142*  

In regions that contain hot springs, volcanic activity, or similar hot and wet conditions, firenewts might be found. These humanoid amphibians live in a militaristic theocracy that reveres elemental fire in its worst incarnation.

## Heat Seekers

Firenewts need hot water to live and breed. A firenewt becomes sluggish, mentally and physically, after spending a week away from an external source of moist heat. A prolonged lack of heat can shut down a firenewt community, as the creatures within go into hibernation and their eggs stop developing.

Firenewts delve for sources of heat in the earth, such as boiling mud and hot springs, that make ideal places to settle. Through excavation and mining in the area, they fashion living space and obtain an ample supply of minerals for other uses, such as smelting, smithing, and alchemy. A firenewt lair features a network of channels and sluices to circulate hot liquid through the settlement.

The alchemy practiced by firenewts focuses on fire. One of their favorite mixtures is a paste of sulfur, mineral salts, and oil. Firenewts chew this blend habitually, because doing so produces a pleasant internal heat and it enables a firenewt to vomit forth a small ball of flame. Most firenewts carry a container with this mixture in it.

Warlocks of Imix command warriors to prove their worth by going on raids to bring back treasure and captives. The warlocks take the choicest loot as a tithe to Imix, and then those who participated in the raid divide the rest according to merit. Prisoners that have no apparent usefulness are sacrificed to Imix and then eaten. Those that are deemed capable of mining and performing other chores around the lair are kept as slaves for a while before meeting the same fate.

When firenewts muster for war, rather than merely staging occasional raids, they take no prisoners. Their goal is nothing less than the annihilation of their foes—and they reserve their greatest animosity for others of their kind. If two groups of firenewts come upon each other, it's likely that they're in competition for the same territory, and a bloody battle is the usual result.

## Religious Militants

Firenewt society and culture are based on the worship of Imix, the Prince of Evil Fire. This veneration of Imix leads firenewts to be aggressive, wrathful, and cruel. Firenewt warlocks of Imix teach that by demonstrating these qualities, a firenewt warrior in combat can become "touched by the Fire Lord," entering a nearly unstoppable battle rage.

## Giant Striders

Firenewts have a close relationship with a type of monstrous beast they believe Imix sent to aid them-borne out by the creatures' ability to send a gout of flame against distant enemies. Called giant striders, these monsters appear birdlike and reptilian, but are truly neither. Firenewts provide shelter, food, and breeding grounds in their lairs for giant striders, and the striders voluntarily serve as mounts for elite fire newt soldiers.

```statblock
"name": "Firenewt Warrior (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "firenewt"
"alignment": "Neutral Evil"
"ac": !!int "16"
"ac_class": "[chain shirt](/3-Mechanics/CLI/items/chain-shirt.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "22"
"hit_dice": "4d8 + 4"
"stats":
- !!int "10"
- !!int "13"
- !!int "12"
- !!int "7"
- !!int "11"
- !!int "8"
"speed": "30 ft."
"damage_immunities": "fire"
"senses": "passive Perception 10"
"languages": "Draconic, Ignan"
"cr": "1/2"
"traits":
- "desc": "The firenewt can breathe air and water."
  "name": "Amphibious"
"actions":
- "desc": "The firenewt makes two attacks with its scimitar."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 1|text(4) (1d6 + 1) slashing damage."
  "name": "Scimitar"
- "desc": "The firenewt spits fire at a creature within 10 feet of it. The creature\
    \ must make a DC 11 Dexterity saving throw, taking dice:2d8|text(9) (2d8)\
    \ fire damage on a failed save, or half as much damage on a successful one"
  "name": "Spit Fire (Recharges after a Short or Long Rest)"
"source":
- "VGM"
- "ToA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Firenewt%20Warrior.webp"
```
^statblock

```encounter-table
name: Firenewt Warrior
creatures:
 - 1: Firenewt Warrior
```

## Environment

underdark, mountain, hill, desert