package kit

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/enchantment"
	"github.com/df-mc/dragonfly/server/item/potion"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type NoDebuff struct {
	Kit
}

func (k *NoDebuff) Name() string {
	return "NoDebuff"
}

func (k *NoDebuff) Armor() []item.Stack {
	items := [4]item.Stack{
		item.NewStack(item.Helmet{Tier: item.ArmourTierDiamond{}}, 1),
		item.NewStack(item.Chestplate{Tier: item.ArmourTierDiamond{}}, 1),
		item.NewStack(item.Leggings{Tier: item.ArmourTierDiamond{}}, 1),
		item.NewStack(item.Boots{Tier: item.ArmourTierDiamond{}}, 1),
	}

	enchantments := [1]item.Enchantment{
		item.NewEnchantment(enchantment.Unbreaking{}, 3),
	}

	for i := 0; i < len(items); i++ {
		for j := 0; j < len(enchantments); j++ {
			items[i] = items[i].WithEnchantments(enchantments[j])
		}
	}

	return items[:]
}

func (k *NoDebuff) Items() []item.Stack {
	items := make([]item.Stack, 36)

	items[0] = item.NewStack(item.Sword{Tier: item.ToolTierDiamond}, 1)
	items[1] = item.NewStack(item.EnderPearl{}, 16)

	for i := 2; i < 36; i++ {
		items[i] = item.NewStack(item.SplashPotion{Type: potion.StrongHealing()}, 1)
	}

	return items[:]
}

func (k *NoDebuff) Apply(p *player.Player) {
	p.Inventory().Clear()
	p.Armour().Clear()
	p.SetGameMode(world.GameModeSurvival)

	armorInv := p.Armour().Inventory()
	for slot, stack := range k.Armor() {

		err := armorInv.SetItem(slot, stack.WithCustomName("Suki"))
		if err != nil {
			panic(err.Error())
			return
		}
	}

	for index, stack := range k.Items() {
		err := p.Inventory().SetItem(index, stack.WithCustomName("Suki"))
		if err != nil {
			panic(err.Error())
			return
		}
	}
}
