include $(TOPDIR)/rules.mk

PKG_NAME:=zimacube-metrics
PKG_VERSION:=0.1
PKG_RELEASE:=1

PKG_SOURCE:=$(PKG_NAME).raw
PKG_URL:=https://github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/releases/download/v0.0.2/zimacube-metrics.raw

define Package/zimacube-metrics
	SECTION:=base
	CATEGORY:=Utilities
	TITLE:=zimacube-metrics
endef

define Package/zimacube-metrics/description
	zimacube-metrics
endef

define Package/zimacube-metrics/install
	$(INSTALL_DIR) $(1)/var/lib/extensions
	$(INSTALL_BIN) ./files/zimacube-metrics $(1)/var/lib/extensions/zimacube-metrics.raw
endef

$(eval $(call BuildPackage,zimacube-metrics))
