---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1
- monster/environment/desert
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/firenewt
aliases: ["Firenewt Warlock of Imix"]
NoteIcon: monster
BestiaryType: humanoid (firenewt)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 143
---
# [Firenewt Warlock of Imix](3-Mechanics\CLI\bestiary\humanoid/firenewt-warlock-of-imix-vgm.md)
*Source: Volo's Guide to Monsters p. 143*  

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
"name": "Firenewt Warlock of Imix (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "firenewt"
"alignment": "Neutral Evil"
"ac": !!int "10"
"ac_class": "13 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "33"
"hit_dice": "6d8 + 6"
"stats":
- !!int "13"
- !!int "11"
- !!int "12"
- !!int "9"
- !!int "11"
- !!int "14"
"speed": "30 ft."
"damage_immunities": "fire"
"senses": "darkvision 120 ft. (penetrates magical darkness), passive Perception 10"
"languages": "Draconic, Ignan"
"cr": "1"
"traits":
- "desc": "The firenewt's innate spellcasting ability is Charisma. It can innately\
    \ cast [mage armor](/3-Mechanics/CLI/spells/mage-armor.md) (self only) at will,\
    \ requiring no material components.\n\nAt will: [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
  "name": "Innate Spellcasting"
- "desc": "The firenewt is a 3rd-level spellcaster. Its spellcasting ability is Charisma\
    \ (spell save DC 12, dice: d20+4 (+4) to hit with spell attacks). It regains\
    \ its expended spell slots when it finishes a short or long rest. It knows the\
    \ following warlock spells:\n\nCantrips (at will): [fire bolt](/3-Mechanics/CLI/spells/fire-bolt.md),\
    \ [guidance](/3-Mechanics/CLI/spells/guidance.md), [light](/3-Mechanics/CLI/spells/light.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\
    \n1st-2nd level (2 slots): [burning hands](/3-Mechanics/CLI/spells/burning-hands.md),\
    \ [flaming sphere](/3-Mechanics/CLI/spells/flaming-sphere.md), [hellish rebuke](/3-Mechanics/CLI/spells/hellish-rebuke.md),\
    \ [scorching ray](/3-Mechanics/CLI/spells/scorching-ray.md)"
  "name": "Spellcasting"
- "desc": "The firenewt can breathe air and water."
  "name": "Amphibious"
- "desc": "When the firenewt reduces an enemy to 0 hit points, the firenewt gains\
    \ 5 temporary hit points."
  "name": "Imix's Blessing"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 1|text(5) (1d8 + 1) piercing damage."
  "name": "Morningstar"
"source":
- "VGM"
- "ToA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Firenewt%20Warlock%20of%20Imix.webp"
```
^statblock

```encounter-table
name: Firenewt Warlock of Imix
creatures:
 - 1: Firenewt Warlock of Imix
```

## Environment

underdark, mountain, hill, desert