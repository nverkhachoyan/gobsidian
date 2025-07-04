---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/11
- monster/environment/desert
- monster/size/large
- monster/type/elemental
aliases: ["Efreeti"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Monster Manual p. 145, Tasha's Cauldron of Everything. Available in the SRD.
---
# [Efreeti](3-Mechanics\CLI\bestiary\elemental/efreeti.md)
*Source: Monster Manual p. 145, Tasha's Cauldron of Everything. Available in the SRD.*  

Hulking genies of the Elemental Plane of Fire, the efreet are masters of flame, immune to fire and able to create it on a whim. Fine silk caftans and damask robes drape their magma-red or coal-black skin, and they bedeck themselves in brass and gold torcs, chains, and rings, all glittering with jewels. When an efreeti flies, its lower body transforms into a column of smoke and embers.

## Haughty and Cruel

The efreet are deceptive, cunning, and cruel to the point of ruthlessness. They despise being forced into servitude and are relentless in pursuit of vengeance against creatures that have wronged them. Efreet don't see themselves in this light, naturally, and regard their race as fair and orderly, even as they admit to an enlightened sense of self-interest.

## Spiteful Slavers

Efreet view all other creatures as enemies or potential serfs. They raid the Material Plane and the elemental planes for slaves, which they capture and bring back to their homes on the Elemental Plane of Fire. The efreet rule as oppressive tyrants, promoting only the cruelest among their slaves. Those overseers are given whips to help keep the rank-and-file slaves in line.

## Planar Raiders

Most efreet reside on the Elemental Plane of Fire, either in great domed fortresses of black glass and basalt surrounded by churning lakes of fire, or in the fabled City of Brass. Additionally, efreet military outposts thronging with their minions and slaves can be found scattered throughout the planes.

On the Material Plane, efreet dwell in fiery regions such as volcanoes and the burning expanses of the world's deserts. Their love of the desert brings them into conflict with the djinn that ride the desert whirlwinds, and with the earthbound dao. Efreet utterly despise marids, with whom they have maintained a passionate conflict throughout the history of both races.

> [!quote]- A quote from A djinni enticing adventurers to free her caliph from a magic lamp in the Charcoal Palace of the City of Brass  
> 
> The armies of the Grand Sultan are bolstered by legions of devils, his palace warded by spells of a thousand archmagi. No one has plundered the efreeti's fabled vaults and lived to tell the tale. By the grace of a thousand winds, you could be the first.

## Genies

Genies are rare elemental creatures out of story and legend. Only a few can be found on the Material Plane. The rest reside on the Elemental Planes, where they rule from lavish palaces and are attended by worshipful slaves.

Genies are as brilliant as they are mighty, as proud as they are majestic. Haughty and decadent, they have a profound sense of entitlement that stems from the knowledge that few creatures except the gods and other genies can challenge their power.

### Creatures of the Elements

A genie is born when the soul of a sentient living creature melds with the primordial matter of an elemental plane. Only under rare circumstances does such an elemental-infused soul coalesce into a manifest form and create a genie.

A genie usually retains no connection to the soul that gave it form. That life force is a building block that determines the genie's form and apparent gender, as well as one or two key personality traits. Although they resemble humanoid beings, genies are elemental spirits given physical form. They don't mate with other genies or produce genie offspring, as all new genies are born out of the same mysterious fusion of spirit energy and elemental power. A genie with a stronger connection to its mortal soul might choose to sire a child with a mortal, although such offspring are rare.

When a genie perishes, it leaves nothing behind except what it was wearing or carrying, along with a small trace of its native element: a pile of dust, a gust of wind, a flash of fire and smoke, or a burst of water and foam.

In contrast to their love of slaves, most genies loathe being bound to service themselves. A genie obeys the will of another only when bribed or compelled by magic. All genies command the power of their native element, but a rare few also possess the power to grant wishes. For both these reasons, mortal mages often seek to bind genies into service.

Noble genies cultivate the jealousy and envy of other genies, asserting their superiority at every opportunity. Other genies respect the influence of the noble genies, knowing how unwise it is to defy a creature that can alter reality at a whim. A genie isn't beholden to any noble genie, however, and will sometimes choose to defy a noble genie's will and risk the consequences.

Their miraculous powers, the grandeur of their abodes, and the numbers of their slaves allow some genies to deceive themselves into believing they are as powerful as the gods. Some go so far as to demand that mortals of other realms-even whole continents or worlds-bow down before them.

### Rule or Be Ruled

Mortal slaves serve to validate a genie's power and high self-opinion. A hundred flattering voices are music to a genie's ears, while two hundred mortal slaves prostrated at its feet are proof that it is lord and master. Genies view slaves as living property, and a genie without property amounts to nothing among its own kind. As a result, many genies treasure their slaves, treating them as honored members of their households. Evil genies freely threaten and abuse their slaves, but never to the extent that the slaves are no longer of use.

### Decadent Nobility

Noble genies are the rarest of their kind. They are used to getting what they want, and have learned to trade their ability to grant wishes to attain the objects of their desire. This constant indulgence has made them decadent, while their supreme power over reality makes them haughty and arrogant. Their vast palaces overflow with wonders and sensory delights beyond imagination.

### The Power of Worship

Genies acknowledge the gods as powerful entities but have no desire to court or worship them. They find the endless fawning and mewling of religious devotees tiresome-except as it is directed toward them by their worshipful slaves.

```statblock
"name": "Efreeti"
"size": "Large"
"type": "elemental"
"alignment": "Lawful Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "200"
"hit_dice": "16d10 + 112"
"stats":
- !!int "22"
- !!int "12"
- !!int "24"
- !!int "16"
- !!int "15"
- !!int "16"
"speed": "40 ft., fly 60 ft."
"saves":
  "Charisma": !!int "7"
  "Wisdom": !!int "6"
  "Intelligence": !!int "7"
"damage_immunities": "fire"
"senses": "darkvision 120 ft., passive Perception 12"
"languages": "Ignan"
"cr": "11"
"traits":
- "desc": "The efreeti's innate spellcasting ability is Charisma (spell save DC 15,\
    \ dice: d20+7 (+7) to hit with spell attacks). It can innately cast the following\
    \ spells, requiring no material components:\n\nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md)\n\
    \n1/day each: [conjure elemental](/3-Mechanics/CLI/spells/conjure-elemental.md)\
    \ ([fire elemental](/3-Mechanics/CLI/bestiary/elemental/fire-elemental.md) only),\
    \ [gaseous form](/3-Mechanics/CLI/spells/gaseous-form.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [major image](/3-Mechanics/CLI/spells/major-image.md), [plane shift](/3-Mechanics/CLI/spells/plane-shift.md),\
    \ [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)\n\n3/day each: [enlarge/reduce](/3-Mechanics/CLI/spells/enlarge-reduce.md),\
    \ [tongues](/3-Mechanics/CLI/spells/tongues.md)"
  "name": "Innate Spellcasting"
- "desc": "If the efreeti dies, its body disintegrates in a flash of fire and puff\
    \ of smoke, leaving behind only equipment the efreeti was wearing or carrying."
  "name": "Elemental Demise"
"actions":
- "desc": "The efreeti makes two scimitar attacks or uses its Hurl Flame twice."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 6|text(13) (2d6 + 6) slashing damage plus dice:2d6|text(7)\
    \ (2d6) fire damage."
  "name": "Scimitar"
- "desc": "Ranged Spell Attack: dice: d20+7 (+7) to hit, range 120 ft., one\
    \ target. Hit: dice:5d6|text(17) (5d6) fire damage."
  "name": "Hurl Flame"
"source":
- "MM"
- "PotA"
- "RoT"
- "TftYP"
- "WDMM"
- "EGW"
- "TCE"
- "JttRC"
- "KftGV"
- "SatO"
- "DoDk"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Efreeti.webp"
```
^statblock

```encounter-table
name: Efreeti
creatures:
 - 1: Efreeti
```

## Environment

desert