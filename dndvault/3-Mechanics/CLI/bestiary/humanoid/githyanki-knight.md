---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/8
- monster/size/medium
- monster/type/humanoid/gith
aliases: ["Githyanki Knight"]
NoteIcon: monster
BestiaryType: humanoid (gith)
SourceType: Bestiary
BookSource: Monster Manual p. 160
---
# [Githyanki Knight](3-Mechanics\CLI\bestiary\humanoid/githyanki-knight.md)
*Source: Monster Manual p. 160*  

The githyanki plunder countless worlds from the decks of their astral vessels and the backs of red dragons. Feathers, beads, gems, and precious metals decorate their armor and weapons-the legendary silver swords with which they cut through their foes. Since winning their freedom from the mind flayers, the githyanki have become ruthless conquerors under the rulership of their dread lich-queen, Vlaakith.

## Astral Raiders

The githyanki despise all other races, undertaking devastating raids that take them from their strongholds in the Astral Plane to the far flung corners of the multiverse. War is the ultimate expression of githyanki culture, and their pitiless black eyes know no mercy. After a raid, they leave shattered survivors enough food and resources to weakly endure. Later, the githyanki return to their conquered foes, plundering them again and again.

## Followers of Gith

In their own language, githyanki means "followers of Gith." Under the guidance of Gith, the githyanki stratified into a militaristic society, with a strict caste system, dedicated to the ongoing fight against the victims and sworn enemies of their race. When their leader Gith perished, she was replaced by her undead adviser, Vlaakith. The lich-queen forbade worship of all beings except herself.

Of all their enemies, the githyanki most hate their former masters, the mind flayers. Their close kin, the githzerai, are second in their enmity. All other creatures are treated with simple contempt by the githyanki, whose xenophobic pride defines their view of inferior races.

## Silver Swords

In ancient times, gith knights created special weapons to combat their mind flayer masters.

These silver swords channel the force of the wielder's will, dealing psychic as well as physical damage. A githyanki can't become a knight until it masters the singular discipline needed to will such a blade into existence. A silver sword is equivalent to a [greatsword](/3-Mechanics/CLI/items/greatsword.md), and takes on the properties of a [+3 greatsword](/3-Mechanics/CLI/items/3-weapon.md) in the hands of its creator.

In the eyes of the githyanki, each silver sword is a priceless relic and a work of art. Githyanki knights will hunt down and destroy any non-githyanki that dares to carry or wield a silver sword, reclaiming it for their people.

## Red Dragon Riders

In the uprising against the illithids, Gith sought allies. Her adviser Vlaakith appealed to Tiamat, the goddess of evil dragonkind, and Gith ventured into the Nine Hells to meet with her. Only Tiamat now knows what passed between them, but Vlaakith returned to the Astral Plane with the Dragon Queen's red dragon consort Ephelomon, who proclaimed that his kind would forever act as allies to the githyanki. Not all red dragons honor the alliance kindled so long ago, but most at least don't consider the githyanki their enemies.

### Outposts in the Mortal Realm

Since creatures that dwell on the Astral Plane don't age, the githyanki establish creches in remote areas of the Material Plane to raise their young. Doubling as military academies, these creches train young githyanki to harness their psychic and combat abilities. When a githyanki grows to adulthood and slays a mind flayer as part of a sacred rite of passage, it is permitted to rejoin its people on the Astral Plane.

## Gith

The warlike githyanki and the contemplative githzerai are a sundered people-two cultures that utterly despise one another. Before there were githyanki or githzerai, these creatures were a single race enslaved by the [mind flayers](/3-Mechanics/CLI/bestiary/aberration/mind-flayer.md). Although they attempted to overthrow their masters many times, their rebellions were repeatedly crushed until a great leader named Gith arose.

After much bloodshed, Gith and her followers threw off the yoke of their illithid masters, but another leader named Zerthimon emerged in the aftermath of battle.

Zerthimon challenged Gith's motives, claiming that her strict martial leadership and desire for vengeance amounted to little more than another form of slavery for her people. A rift erupted between followers of each leader, and they eventually became the two races whose enmity endures to this day.

Whether these tall, gaunt creatures were peaceful or savage, cultured or primitive before the [mind flayers](/3-Mechanics/CLI/bestiary/aberration/mind-flayer.md) enslaved and changed them, none can say. Not even the original name of their race remains from that distant time.

> [!quote]- A quote from Aristul the Yellow, master of planar lore  
> 
> The githyanki and the githzerai were so profoundly scarred by their enslavement to the mind flayers that they forget they were one race, united. Having won their freedom, they wage war against each other with a hatred none can fully comprehend.


```statblock
"name": "Githyanki Knight"
"size": "Medium"
"type": "humanoid"
"subtype": "gith"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "91"
"hit_dice": "14d8 + 28"
"stats":
- !!int "16"
- !!int "14"
- !!int "15"
- !!int "14"
- !!int "14"
- !!int "15"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "5"
  "Intelligence": !!int "5"
  "Constitution": !!int "5"
"senses": "passive Perception 12"
"languages": "Gith"
"cr": "8"
"traits":
- "desc": "The githyanki's innate spellcasting ability is Intelligence (spell save\
    \ DC 13, dice: d20+5 (+5) to hit with spell attacks). It can innately cast\
    \ the following spells, requiring no components:\n\nAt will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md)\
    \ (the hand is invisible)\n\n1/day each: [plane shift](/3-Mechanics/CLI/spells/plane-shift.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)\n\n3/day each: [jump](/3-Mechanics/CLI/spells/jump.md),\
    \ [misty step](/3-Mechanics/CLI/spells/misty-step.md), [nondetection](/3-Mechanics/CLI/spells/nondetection.md)\
    \ (self only), [tongues](/3-Mechanics/CLI/spells/tongues.md)"
  "name": "Innate Spellcasting (Psionics)"
"actions":
- "desc": "The githyanki makes two silver greatsword attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 6|text(13) (2d6 + 6) slashing damage plus dice:3d6|text(10)\
    \ (3d6) psychic damage. This is a magic weapon attack. On a critical hit against\
    \ a target in an astral body (as with the [astral projection](/3-Mechanics/CLI/spells/astral-projection.md)\
    \ spell), the githyanki can cut the silvery cord that tethers the target to its\
    \ material body, instead of dealing damage."
  "name": "Silver Greatsword"
"source":
- "MM"
- "WDMM"
- "LoX"
- "PaBTSO"
- "AATM"
- "SatO"
- "BMT"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Githyanki%20Knight.webp"
```
^statblock

```encounter-table
name: Githyanki Knight
creatures:
 - 1: Githyanki Knight
```