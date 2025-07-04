---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/12
- monster/environment/coastal
- monster/environment/desert
- monster/environment/grassland
- monster/environment/mountain
- monster/size/huge
- monster/type/celestial
aliases: ["Ki-rin"]
NoteIcon: monster
BestiaryType: celestial
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 163
---
# [Ki-rin](3-Mechanics\CLI\bestiary\celestial/ki-rin-vgm.md)
*Source: Volo's Guide to Monsters p. 163*  

Ki-rins are noble, celestial creatures. In the Outer Planes, ki-rins in service to benevolent deities take a direct role in the eternal struggle between good and evil. In the mortal world, a ki-rin is celebrated far and wide as a harbinger of destiny, a guardian of the sacred, and a counterbalance to the forces of evil.

## Good Personified

Ki-rins are the embodiment of good, and simply beholding one can evoke fear or awe in an observer. A typical ki-rin looks like a muscular stag the size of an elephant, covered in golden scales lined in some places with golden fur. It has a dark gold mane and tail, coppery cloven hooves, and a spiral-shaped coppery horn just above and between its luminous violet eyes. In a breeze or when aloft, the creature's scales and hair can create the impression that the ki-rin is ablaze with a holy, golden fire.

Beyond their coloration, ki-rins vary in appearance, based on the deity each one reveres and the function it typically performs in service to that god. Some are horse-shaped, looking like gigantic unicorns, and are often used as guardians. Others have draconic features and tend to be aggressive foes of evil. One horn is most common, but a ki-rin of fierce demeanor might have two horns or a set of antlers like those of a great stag.

A ki-rin in the world claims a territory to watch over, and one ki-rin might safeguard an area that encompasses several nations. On other planes, ki-rins that serve good deities go wherever they are commanded, which could include coming to the Material Plane on a mission. A ki-rin disciple in the world usually serves its deity as a scout, a messenger, or a spy.

Ki-rins are attracted to the worship of deities of courage, loyalty, selflessness, and truth, as well as the advancement of just societies. For instance, in Faerûn, ki-rins rally mostly to Torm, although ki-rins also serve his allies Tyr and Ilmater.

Many who seek a ki-rin's guidance end up pledging service to the creature. They study as monks under its tutelage and serve as its agents in the world. The followers of a ki-rin might travel incognito across the land, seeking news of growing evil and working behind the scenes, or might be champions of their master's cause, out to defeat villainy wherever it is found.

## Bringers of Boons

Common folk consider ki-rins to be rare and remote heralds of good fortune. Seeing a ki-rin fly overhead is a blessing, and events that happen on such a day are especially auspicious. If a ki-rin alights during a ceremony, such as a birth announcement or a coronation, everyone present understands that the creature is telling them great good could be in the offing. The ki-rin conveys its gifts and omens, then rises back into the sky. Ki-rins have also been known to appear at the sites of great battles to inspire and strengthen the side of good, or to rescue heroes from certain death.

## Objects of Adoration

Because a ki-rin is renowned for its wisdom, other creatures would naturally seek it out with questions and requests if they could. For that reason among others, the creature makes its lair atop a forbidding mountain peak or in some other equally inaccessible location. Only those that have the tenacity to complete the daunting journey to a ki-rin's lair can prove themselves worthy of speaking with its occupant.

## Lair of Luxury

On the celestial planes, ki-rins reside in lofty, elegant aeries filled with luxurious objects. In the world, a ki-rin chooses a similar location, such as atop a tall pinnacle or within a cloud solidified by the ki-rin's magic. When viewed from the outside, a ki-rin's lair is indistinguishable from a natural site, and the entrance is difficult for visitors to find and reach. Inside, the lair is a serene and comfortable place, its ambiance a mix between palace and temple. If the ki-rin has taken creatures into its service, its lair doubles as a sacred site wherein the ki-rin not only rests, but also teaches of holy mysteries.

```statblock
"name": "Ki-rin (VGM)"
"size": "Huge"
"type": "celestial"
"alignment": "Lawful Good"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "152"
"hit_dice": "16d12 + 48"
"stats":
- !!int "21"
- !!int "16"
- !!int "16"
- !!int "19"
- !!int "20"
- !!int "20"
"speed": "60 ft., fly 120 ft. (hover)"
"skillsaves":
  "Religion": !!int "8"
  "Insight": !!int "9"
  "Perception": !!int "9"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 30 ft., darkvision 120 ft., passive Perception 19"
"languages": "all, telepathy 120 ft."
"cr": "12"
"traits":
- "desc": "The ki-rin's innate spellcasting ability is Charisma (spell save DC 17).\
    \ The ki-rin can innately cast the following spells, requiring no material components:\n\
    \nAt will: [gaseous form](/3-Mechanics/CLI/spells/gaseous-form.md), [major\
    \ image](/3-Mechanics/CLI/spells/major-image.md) (6th-level version), [wind walk](/3-Mechanics/CLI/spells/wind-walk.md)\n\
    \n1/day: [create food and water](/3-Mechanics/CLI/spells/create-food-and-water.md)"
  "name": "Innate Spellcasting"
- "desc": "The ki-rin is a 18th-level spellcaster. Its spellcasting ability is Wisdom\
    \ (spell save DC 17, dice: d20+9 (+9) to hit with spell attacks). It has the\
    \ following cleric spells prepared:\n\nCantrips (at will): [light](/3-Mechanics/CLI/spells/light.md),\
    \ [mending](/3-Mechanics/CLI/spells/mending.md), [sacred flame](/3-Mechanics/CLI/spells/sacred-flame.md),\
    \ [spare the dying](/3-Mechanics/CLI/spells/spare-the-dying.md), [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\
    \n1st level (4 slots): [command](/3-Mechanics/CLI/spells/command.md), [cure\
    \ wounds](/3-Mechanics/CLI/spells/cure-wounds.md), [detect evil and good](/3-Mechanics/CLI/spells/detect-evil-and-good.md),\
    \ [protection from evil and good](/3-Mechanics/CLI/spells/protection-from-evil-and-good.md),\
    \ [sanctuary](/3-Mechanics/CLI/spells/sanctuary.md)\n\n2nd level (3 slots):\
    \ [calm emotions](/3-Mechanics/CLI/spells/calm-emotions.md), [lesser restoration](/3-Mechanics/CLI/spells/lesser-restoration.md),\
    \ [silence](/3-Mechanics/CLI/spells/silence.md)\n\n3rd level (3 slots): [dispel\
    \ magic](/3-Mechanics/CLI/spells/dispel-magic.md), [remove curse](/3-Mechanics/CLI/spells/remove-curse.md),\
    \ [sending](/3-Mechanics/CLI/spells/sending.md)\n\n4th level (3 slots): [banishment](/3-Mechanics/CLI/spells/banishment.md),\
    \ [freedom of movement](/3-Mechanics/CLI/spells/freedom-of-movement.md), [guardian\
    \ of faith](/3-Mechanics/CLI/spells/guardian-of-faith.md)\n\n5th level (3 slots):\
    \ [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md), [mass\
    \ cure wounds](/3-Mechanics/CLI/spells/mass-cure-wounds.md), [scrying](/3-Mechanics/CLI/spells/scrying.md)\n\
    \n6th level (1 slots): [heroes' feast](/3-Mechanics/CLI/spells/heroes-feast.md),\
    \ [true seeing](/3-Mechanics/CLI/spells/true-seeing.md)\n\n7th level (1 slots):\
    \ [etherealness](/3-Mechanics/CLI/spells/etherealness.md), [plane shift](/3-Mechanics/CLI/spells/plane-shift.md)\n\
    \n8th level (1 slots): [control weather](/3-Mechanics/CLI/spells/control-weather.md)\n\
    \n9th level (1 slots): [true resurrection](/3-Mechanics/CLI/spells/true-resurrection.md)"
  "name": "Spellcasting"
- "desc": "If the ki-rin fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The ki-rin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The ki-rin's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The ki-rin makes three attacks: two with its hooves and one with its horn."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 15 ft., one target.\
    \ Hit: dice:2d4 + 5|text(10) (2d4 + 5) bludgeoning damage."
  "name": "Hoof"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 5|text(14) (2d8 + 5) piercing damage."
  "name": "Horn"
"legendary_actions":
- "desc": "The ki-rin makes a Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ check or a Wisdom ([Insight](/3-Mechanics/CLI/rules/skills.md#Insight)) check."
  "name": "Detect"
- "desc": "The ki-rin makes a hoof attack or casts [sacred flame](/3-Mechanics/CLI/spells/sacred-flame.md)."
  "name": "Smite"
- "desc": "The ki-rin moves up to its half its speed without provoking opportunity\
    \ attacks."
  "name": "Move"
"lair_actions":
- "desc": "Inside its lair, a ki-rin has the power to conjure objects up to three\
    \ times per day, using each of the following versions of the power once. One version\
    \ permanently creates enough objects made of soft, plant-based material - including\
    \ manufactured objects such as cloth, pillows, rope, blankets, and clothing -\
    \ to fill a cube 20 feet on a side. The second version permanently creates enough\
    \ objects made of wood, or similarly hard plant-based material, to fill a cube\
    \ 10 feet on a side. The third version creates enough objects made of stone or\
    \ metal to fill a cube 2 feet on a side, but any materials created in this way\
    \ last for only 1 hour."
  "name": ""
"regional_effects":
- "desc": "The ki-rin's celestial nature transforms the region around its lair. Any\
    \ of the following magical effects is possible for travelers to encounter in the\
    \ vicinity:"
  "name": ""
- "desc": "- Water flows pure within 3 miles of a ki-rin's lair. Any purposeful corruption\
    \ of the water lasts for no longer than 3 minutes.  \n- Animals, plants, and good\
    \ creatures within 3 miles of the ki-rin's lair gain vigor as they evolve toward\
    \ an idealized form. Such creatures are rarely aggressive toward others that aren't\
    \ normally prey. Evil creatures can't tolerate the holy atmosphere within the\
    \ same distance, and usually choose to live much farther from the domain of a\
    \ ki-rin.  \n- Curses, diseases, and poisons affecting good-aligned creatures\
    \ are suppressed when those creatures are within 3 miles of the lair.  \n- A ki-rin\
    \ can cast [control weather](/3-Mechanics/CLI/spells/control-weather.md) while\
    \ it is within 3 miles of its lair. The spell's point of origin is always the\
    \ point outdoors closest to the center of its lair. The ki-rin doesn't need to\
    \ maintain a clear path to the sky or to concentrate for the change in weather\
    \ to persist.  \n- Within 3 miles of the lair, winds buoy non-evil creatures that\
    \ fall due to no act of the ki-rin or its allies. Such creatures descend at a\
    \ rate of 60 feet per round and take no falling damage.  "
  "name": ""
- "desc": "When the ki-rin dies, all these effects disappear immediately, although\
    \ the invigorating effect on flora and fauna remains for 3 years."
  "name": ""
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Ki-rin.webp"
```
^statblock

```encounter-table
name: Ki-rin
creatures:
 - 1: Ki-rin
```

## Environment

mountain, grassland, desert, coastal