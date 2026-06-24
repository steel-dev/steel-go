package steel

import (
	"encoding/json"
	"github.com/steel-dev/steel-go/internal/apijson"
	"github.com/steel-dev/steel-go/internal/param"
	"time"
)

type ProfileStatus string

const (
	ProfileStatusUploading ProfileStatus = "UPLOADING"
	ProfileStatusReady     ProfileStatus = "READY"
	ProfileStatusFailed    ProfileStatus = "FAILED"
)

type ScrapeRequestFormatItem string

const (
	ScrapeRequestFormatItemHTML        ScrapeRequestFormatItem = "html"
	ScrapeRequestFormatItemReadability ScrapeRequestFormatItem = "readability"
	ScrapeRequestFormatItemCleanedHTML ScrapeRequestFormatItem = "cleaned_html"
	ScrapeRequestFormatItemMarkdown    ScrapeRequestFormatItem = "markdown"
)

type SessionCostResponseCurrency string

const (
	SessionCostResponseCurrencyUsd SessionCostResponseCurrency = "usd"
)

type SessionCostResponseUnit string

const (
	SessionCostResponseUnitDollar SessionCostResponseUnit = "dollar"
)

type SessionResponseProxySource string

const (
	SessionResponseProxySourceSteel    SessionResponseProxySource = "steel"
	SessionResponseProxySourceExternal SessionResponseProxySource = "external"
)

type SessionResponseRegion string

const (
	SessionResponseRegionLax         SessionResponseRegion = "lax"
	SessionResponseRegionOrd         SessionResponseRegion = "ord"
	SessionResponseRegionIad         SessionResponseRegion = "iad"
	SessionResponseRegionScl         SessionResponseRegion = "scl"
	SessionResponseRegionFra         SessionResponseRegion = "fra"
	SessionResponseRegionNrt         SessionResponseRegion = "nrt"
	SessionResponseRegionUsEast      SessionResponseRegion = "us-east"
	SessionResponseRegionUsWest      SessionResponseRegion = "us-west"
	SessionResponseRegionUsCentral   SessionResponseRegion = "us-central"
	SessionResponseRegionEuWest      SessionResponseRegion = "eu-west"
	SessionResponseRegionEuCentral   SessionResponseRegion = "eu-central"
	SessionResponseRegionApNortheast SessionResponseRegion = "ap-northeast"
	SessionResponseRegionApSoutheast SessionResponseRegion = "ap-southeast"
	SessionResponseRegionSaEast      SessionResponseRegion = "sa-east"
)

type SessionResponseReleaseReason string

const (
	SessionResponseReleaseReasonUserRequested     SessionResponseReleaseReason = "user_requested"
	SessionResponseReleaseReasonTimeout           SessionResponseReleaseReason = "timeout"
	SessionResponseReleaseReasonInactivityTimeout SessionResponseReleaseReason = "inactivity_timeout"
	SessionResponseReleaseReasonCreationTimeout   SessionResponseReleaseReason = "creation_timeout"
	SessionResponseReleaseReasonStartupFailed     SessionResponseReleaseReason = "startup_failed"
	SessionResponseReleaseReasonBrowserClosed     SessionResponseReleaseReason = "browser_closed"
	SessionResponseReleaseReasonBrowserCrashed    SessionResponseReleaseReason = "browser_crashed"
)

type SessionResponseStatus string

const (
	SessionResponseStatusLive     SessionResponseStatus = "live"
	SessionResponseStatusReleased SessionResponseStatus = "released"
	SessionResponseStatusFailed   SessionResponseStatus = "failed"
)

type ComputerActionRequestVariant0Action string

const (
	ComputerActionRequestVariant0ActionMoveMouse ComputerActionRequestVariant0Action = "move_mouse"
)

type ComputerActionRequestVariant1Action string

const (
	ComputerActionRequestVariant1ActionClickMouse ComputerActionRequestVariant1Action = "click_mouse"
)

type ComputerActionRequestVariant1Button string

const (
	ComputerActionRequestVariant1ButtonLeft    ComputerActionRequestVariant1Button = "left"
	ComputerActionRequestVariant1ButtonRight   ComputerActionRequestVariant1Button = "right"
	ComputerActionRequestVariant1ButtonMiddle  ComputerActionRequestVariant1Button = "middle"
	ComputerActionRequestVariant1ButtonBack    ComputerActionRequestVariant1Button = "back"
	ComputerActionRequestVariant1ButtonForward ComputerActionRequestVariant1Button = "forward"
)

type ComputerActionRequestVariant1ClickType string

const (
	ComputerActionRequestVariant1ClickTypeDown  ComputerActionRequestVariant1ClickType = "down"
	ComputerActionRequestVariant1ClickTypeUp    ComputerActionRequestVariant1ClickType = "up"
	ComputerActionRequestVariant1ClickTypeClick ComputerActionRequestVariant1ClickType = "click"
)

type ComputerActionRequestVariant2Action string

const (
	ComputerActionRequestVariant2ActionDragMouse ComputerActionRequestVariant2Action = "drag_mouse"
)

type ComputerActionRequestVariant3Action string

const (
	ComputerActionRequestVariant3ActionScroll ComputerActionRequestVariant3Action = "scroll"
)

type ComputerActionRequestVariant4Action string

const (
	ComputerActionRequestVariant4ActionPressKey ComputerActionRequestVariant4Action = "press_key"
)

type ComputerActionRequestVariant5Action string

const (
	ComputerActionRequestVariant5ActionTypeText ComputerActionRequestVariant5Action = "type_text"
)

type ComputerActionRequestVariant6Action string

const (
	ComputerActionRequestVariant6ActionWait ComputerActionRequestVariant6Action = "wait"
)

type ComputerActionRequestVariant7Action string

const (
	ComputerActionRequestVariant7ActionTakeScreenshot ComputerActionRequestVariant7Action = "take_screenshot"
)

type ComputerActionRequestVariant8Action string

const (
	ComputerActionRequestVariant8ActionGetCursorPosition ComputerActionRequestVariant8Action = "get_cursor_position"
)

type CreateSessionRequestDeviceConfigDevice string

const (
	CreateSessionRequestDeviceConfigDeviceDesktop CreateSessionRequestDeviceConfigDevice = "desktop"
	CreateSessionRequestDeviceConfigDeviceMobile  CreateSessionRequestDeviceConfigDevice = "mobile"
)

type CreateSessionRequestSessionContextCookiesItemPriority string

const (
	CreateSessionRequestSessionContextCookiesItemPriorityLow    CreateSessionRequestSessionContextCookiesItemPriority = "Low"
	CreateSessionRequestSessionContextCookiesItemPriorityMedium CreateSessionRequestSessionContextCookiesItemPriority = "Medium"
	CreateSessionRequestSessionContextCookiesItemPriorityHigh   CreateSessionRequestSessionContextCookiesItemPriority = "High"
)

type CreateSessionRequestSessionContextCookiesItemSameSite string

const (
	CreateSessionRequestSessionContextCookiesItemSameSiteStrict CreateSessionRequestSessionContextCookiesItemSameSite = "Strict"
	CreateSessionRequestSessionContextCookiesItemSameSiteLax    CreateSessionRequestSessionContextCookiesItemSameSite = "Lax"
	CreateSessionRequestSessionContextCookiesItemSameSiteNone   CreateSessionRequestSessionContextCookiesItemSameSite = "None"
)

type CreateSessionRequestSessionContextCookiesItemSourceScheme string

const (
	CreateSessionRequestSessionContextCookiesItemSourceSchemeUnset     CreateSessionRequestSessionContextCookiesItemSourceScheme = "Unset"
	CreateSessionRequestSessionContextCookiesItemSourceSchemeNonSecure CreateSessionRequestSessionContextCookiesItemSourceScheme = "NonSecure"
	CreateSessionRequestSessionContextCookiesItemSourceSchemeSecure    CreateSessionRequestSessionContextCookiesItemSourceScheme = "Secure"
)

type CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity string

const (
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityACoruna                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "A_CORUNA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAbidjan                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ABIDJAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAbuDhabi                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ABU_DHABI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAbuja                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ABUJA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAcapulcoDeJuarez        CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ACAPULCO_DE JUAREZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAccra                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ACCRA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAdana                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ADANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAdapazari               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ADAPAZARI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAddisAbaba              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ADDIS_ABABA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAdelaide                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ADELAIDE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAfyonkarahisar          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AFYONKARAHISAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAgadir                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AGADIR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAguasLindasDeGoias      CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AGUAS_LINDAS DE GOIAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAguascalientes          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AGUASCALIENTES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAhmedabad               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AHMEDABAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAizawl                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AIZAWL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAjman                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AJMAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAkron                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AKRON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAksaray                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AKSARAY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlAinCity               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AL_AIN CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlMansurah              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AL_MANSURAH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlQatif                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AL_QATIF"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlajuela                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ALAJUELA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlbany                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ALBANY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlbuquerque             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ALBUQUERQUE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlexandria              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ALEXANDRIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlgiers                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ALGIERS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlicante                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ALICANTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlmada                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ALMADA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlmaty                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ALMATY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlmereStad              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ALMERE_STAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAlvorada                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ALVORADA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAmadora                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AMADORA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAmasya                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AMASYA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAmbato                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AMBATO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAmericana               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AMERICANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAmman                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AMMAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAmsterdam               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AMSTERDAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAnanindeua              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANANINDEUA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAnapolis                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANAPOLIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAngelesCity             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANGELES_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAngers                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANGERS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAngraDosReis            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANGRA_DOS REIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAnkara                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANKARA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAntakya                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANTAKYA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAntalya                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANTALYA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAntananarivo            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANTANANARIVO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAntipoloCity            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANTIPOLO_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAntofagasta             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANTOFAGASTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAntwerp                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ANTWERP"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAparecidaDeGoiania      CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "APARECIDA_DE GOIANIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityApodaca                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "APODACA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAracaju                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ARACAJU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAracatuba               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ARACATUBA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityArad                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ARAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAraguaina               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ARAGUAINA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityArapiraca               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ARAPIRACA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAraraquara              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ARARAQUARA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityArequipa                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AREQUIPA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityArica                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ARICA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityArlington               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ARLINGTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAryanah                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ARYANAH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAstana                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ASTANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAsuncion                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ASUNCION"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAsyut                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ASYUT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAtakum                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ATAKUM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAthens                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ATHENS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAtibaia                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ATIBAIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAtlanta                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ATLANTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAuburn                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AUBURN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAuckland                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AUCKLAND"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAurora                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AURORA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAustin                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AUSTIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAvellaneda              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AVELLANEDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAydin                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AYDIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityAzcapotzalco            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "AZCAPOTZALCO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBacolodCity             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BACOLOD_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBacoor                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BACOOR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBaghdad                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BAGHDAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBaguioCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BAGUIO_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBahiaBlanca             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BAHIA_BLANCA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBakersfield             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BAKERSFIELD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBaku                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BAKU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBalikesir               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BALIKESIR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBalikpapan              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BALIKPAPAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBalnearioCamboriu       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BALNEARIO_CAMBORIU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBaltimore               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BALTIMORE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBandarLampung           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BANDAR_LAMPUNG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBandarSeriBegawan       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BANDAR_SERI BEGAWAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBandung                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BANDUNG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBangkok                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BANGKOK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBanjaLuka               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BANJA_LUKA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBanjarmasin             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BANJARMASIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBarcelona               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BARCELONA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBari                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BARI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBarquisimeto            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BARQUISIMETO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBarraMansa              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BARRA_MANSA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBarranquilla            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BARRANQUILLA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBarueri                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BARUERI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBatam                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BATAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBatangas                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BATANGAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBatman                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BATMAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBatnaCity               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BATNA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBatonRouge              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BATON_ROUGE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBatumi                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BATUMI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBauru                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BAURU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBeirut                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BEIRUT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBejaia                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BEJAIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBekasi                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BEKASI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBelem                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BELEM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBelfast                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BELFAST"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBelfordRoxo             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BELFORD_ROXO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBelgrade                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BELGRADE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBeloHorizonte           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BELO_HORIZONTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBengaluru               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BENGALURU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBeniMellal              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BENI_MELLAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBerazategui             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BERAZATEGUI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBern                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BERN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBetim                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BETIM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBharatpur               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BHARATPUR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBhopal                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BHOPAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBhubaneswar             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BHUBANESWAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBialystok               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BIALYSTOK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBienHoa                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BIEN_HOA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBilbao                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BILBAO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBilecik                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BILECIK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBiratnagar              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BIRATNAGAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBirmingham              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BIRMINGHAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBishkek                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BISHKEK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBizerte                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BIZERTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBlida                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BLIDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBloemfontein            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BLOEMFONTEIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBloomington             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BLOOMINGTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBlumenau                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BLUMENAU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBoaVista                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BOA_VISTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBochum                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BOCHUM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBogor                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BOGOR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBogota                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BOGOTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBoise                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BOISE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBoksburg                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BOKSBURG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBologna                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BOLOGNA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBolu                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BOLU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBordeaux                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BORDEAUX"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBoston                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BOSTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBotucatu                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BOTUCATU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBradford                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRADFORD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBraga                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRAGA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBragancaPaulista        CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRAGANCA_PAULISTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBrampton                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRAMPTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBrasilia                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRASILIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBrasov                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRASOV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBratislava              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRATISLAVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBremen                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BREMEN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBrescia                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRESCIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBrest                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BREST"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBridgetown              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRIDGETOWN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBrisbane                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRISBANE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBristol                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRISTOL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBrno                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRNO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBrooklyn                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BROOKLYN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBrussels                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BRUSSELS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBucaramanga             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BUCARAMANGA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBucharest               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BUCHAREST"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBudapest                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BUDAPEST"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBuenosAires             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BUENOS_AIRES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBuffalo                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BUFFALO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBukGu                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BUK_GU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBukhara                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BUKHARA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBurgas                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BURGAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBurnaby                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BURNABY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBursa                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BURSA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityButuan                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BUTUAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityBydgoszcz               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "BYDGOSZCZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCabanatuanCity          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CABANATUAN_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCaboFrio                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CABO_FRIO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCabuyao                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CABUYAO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCachoeiroDeItapemirim   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CACHOEIRO_DE ITAPEMIRIM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCagayanDeOro            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAGAYAN_DE ORO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCagliari                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAGLIARI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCairo                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAIRO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCalamba                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CALAMBA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCalgary                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CALGARY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCaloocanCity            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CALOOCAN_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCamacari                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAMACARI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCamaragibe              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAMARAGIBE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCampeche                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAMPECHE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCampinaGrande           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAMPINA_GRANDE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCampinas                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAMPINAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCampoGrande             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAMPO_GRANDE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCampoLargo              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAMPO_LARGO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCamposDosGoytacazes     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAMPOS_DOS GOYTACAZES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCanTho                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAN_THO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCanoas                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CANOAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCanton                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CANTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCapeTown                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAPE_TOWN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCaracas                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CARACAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCaraguatatuba           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CARAGUATATUBA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCarapicuiba             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CARAPICUIBA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCardiff                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CARDIFF"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCariacica               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CARIACICA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCarmona                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CARMONA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCartagena               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CARTAGENA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCaruaru                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CARUARU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCasablanca              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CASABLANCA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCascavel                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CASCAVEL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCaseros                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CASEROS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCastanhal               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CASTANHAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCastries                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CASTRIES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCatalao                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CATALAO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCatamarca               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CATAMARCA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCatanduva               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CATANDUVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCatania                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CATANIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCaucaia                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAUCAIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCaxiasDoSul             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CAXIAS_DO SUL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCebuCity                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CEBU_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCentral                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CENTRAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCentro                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CENTRO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCenturion               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CENTURION"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChaguanas               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHAGUANAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChandigarh              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHANDIGARH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChandler                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHANDLER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChangHua                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHANG_HUA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChapeco                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHAPECO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCharleston              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHARLESTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCharlotte               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHARLOTTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChelyabinsk             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHELYABINSK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChennai                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHENNAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCherkasy                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHERKASY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChernivtsi              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHERNIVTSI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChia                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChiangMai               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHIANG_MAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChiclayo                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHICLAYO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChihuahuaCity           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHIHUAHUA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChimbote                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHIMBOTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChisinau                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHISINAU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChittagong              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHITTAGONG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityChristchurch            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CHRISTCHURCH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCincinnati              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CINCINNATI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCirebon                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CIREBON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCityOfMuntinlupa        CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CITY_OF MUNTINLUPA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCiudadDelEste           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CIUDAD_DEL ESTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCiudadGuayana           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CIUDAD_GUAYANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCiudadJuarez            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CIUDAD_JUAREZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCiudadNezahualcoyotl    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CIUDAD_NEZAHUALCOYOTL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCiudadObregon           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CIUDAD_OBREGON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCleveland               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CLEVELAND"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityClujNapoca              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CLUJ_NAPOCA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCochabamba              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COCHABAMBA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCoimbatore              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COIMBATORE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCoimbra                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COIMBRA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCologne                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COLOGNE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityColombo                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COLOMBO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityColoradoSprings         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COLORADO_SPRINGS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityColumbia                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COLUMBIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityColumbus                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COLUMBUS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityComodoroRivadavia       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COMODORO_RIVADAVIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityConcepcion              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CONCEPCION"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityConcord                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CONCORD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityConstanta               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CONSTANTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityConstantine             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CONSTANTINE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityContagem                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CONTAGEM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCopenhagen              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COPENHAGEN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCordoba                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CORDOBA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCorrientes              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CORRIENTES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCorum                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CORUM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCotia                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COTIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCoventry                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "COVENTRY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCraiova                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CRAIOVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCriciuma                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CRICIUMA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCroydon                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CROYDON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCuautitlanIzcalli       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CUAUTITLAN_IZCALLI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCucuta                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CUCUTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCuenca                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CUENCA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCuernavaca              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CUERNAVACA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCuiaba                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CUIABA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCuliacan                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CULIACAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCuritiba                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CURITIBA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityCusco                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "CUSCO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDaNang                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DA_NANG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDagupan                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DAGUPAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDakar                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DAKAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDallas                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DALLAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDamietta                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DAMIETTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDammam                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DAMMAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDarEsSalaam             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DAR_ES SALAAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDasmarinas              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DASMARINAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDavaoCity               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DAVAO_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDayton                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DAYTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDebrecen                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DEBRECEN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDecatur                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DECATUR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDehradun                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DEHRADUN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDelhi                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DELHI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDenizli                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DENIZLI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDenpasar                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DENPASAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDenver                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DENVER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDepok                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DEPOK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDerby                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DERBY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDetroit                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DETROIT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDhaka                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DHAKA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDiadema                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DIADEMA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDivinopolis             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DIVINOPOLIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDiyarbakir              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DIYARBAKIR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDjelfa                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DJELFA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDnipro                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DNIPRO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDoha                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DOHA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDortmund                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DORTMUND"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDourados                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DOURADOS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDresden                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DRESDEN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDubai                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DUBAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDublin                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DUBLIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDuezce                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DUEZCE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDuisburg                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DUISBURG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDuqueDeCaxias           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DUQUE_DE CAXIAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDurango                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DURANGO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDurban                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DURBAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityDusseldorf              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "DUSSELDORF"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityEcatepec                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ECATEPEC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityEdinburgh               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "EDINBURGH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityEdirne                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "EDIRNE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityEdmonton                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "EDMONTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityElJadida                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "EL_JADIDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityElPaso                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "EL_PASO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityElazig                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ELAZIG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityEmbu                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "EMBU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityEnsenada                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ENSENADA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityErbil                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ERBIL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityErzurum                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ERZURUM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityEskisehir               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ESKISEHIR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityEspoo                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ESPOO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityEssen                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ESSEN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFaisalabad              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FAISALABAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFayetteville            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FAYETTEVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFazendaRioGrande        CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FAZENDA_RIO GRANDE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFeiraDeSantana          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FEIRA_DE SANTANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFes                     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFlorence                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FLORENCE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFlorencioVarela         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FLORENCIO_VARELA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFlorianopolis           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FLORIANOPOLIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFontana                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FONTANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFormosa                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FORMOSA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFortLauderdale          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FORT_LAUDERDALE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFortWayne               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FORT_WAYNE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFortWorth               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FORT_WORTH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFortaleza               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FORTALEZA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFozDoIguacu             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FOZ_DO IGUACU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFranca                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FRANCA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFranciscoMorato         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FRANCISCO_MORATO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFrancoDaRocha           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FRANCO_DA ROCHA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFrankfurtAmMain         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FRANKFURT_AM MAIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFredericksburg          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FREDERICKSBURG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFresno                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FRESNO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityFunchal                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "FUNCHAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGaborone                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GABORONE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGainesville             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GAINESVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGalati                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GALATI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGangnamGu               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GANGNAM_GU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGaranhuns               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GARANHUNS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGatineau                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GATINEAU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGaziantep               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GAZIANTEP"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGdansk                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GDANSK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGdynia                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GDYNIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGeneralTrias            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GENERAL_TRIAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGeneva                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GENEVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGenoa                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GENOA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGeorgeTown              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GEORGE_TOWN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGeorgetown              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GEORGETOWN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGhaziabad               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GHAZIABAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGhent                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GHENT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGijon                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GIJON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGiresun                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GIRESUN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGiza                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GIZA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGlasgow                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GLASGOW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGlendale                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GLENDALE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGliwice                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GLIWICE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGoiania                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GOIANIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGomel                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GOMEL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGothenburg              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GOTHENBURG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGovernadorValadares     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GOVERNADOR_VALADARES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGoyangSi                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GOYANG_SI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGranada                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GRANADA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGrandRapids             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GRAND_RAPIDS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGravatai                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GRAVATAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGraz                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GRAZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGreensboro              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GREENSBORO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGreenville              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GREENVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGuadalajara             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUADALAJARA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGuadalupe               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUADALUPE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGuangzhou               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUANGZHOU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGuarapuava              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUARAPUAVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGuaratingueta           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUARATINGUETA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGuaruja                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUARUJA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGuarulhos               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUARULHOS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGuatemalaCity           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUATEMALA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGuayaquil               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUAYAQUIL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGujranwala              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUJRANWALA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGurugram                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GURUGRAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGustavoAdolfoMadero     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUSTAVO_ADOLFO MADERO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGuwahati                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GUWAHATI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityGwanakGu                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "GWANAK_GU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHackney                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HACKNEY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHaifa                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HAIFA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHaiphong                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HAIPHONG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHamburg                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HAMBURG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHamilton                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HAMILTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHanoi                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HANOI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHanover                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HANOVER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHarare                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HARARE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHavana                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HAVANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHelsinki                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HELSINKI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHenderson               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HENDERSON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHeredia                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HEREDIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHermosillo              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HERMOSILLO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHialeah                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HIALEAH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHoChiMinhCity           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HO_CHI MINH CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHollywood               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HOLLYWOOD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHolon                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HOLON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHonolulu                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HONOLULU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHortolandia             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HORTOLANDIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHrodna                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HRODNA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHsinchu                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HSINCHU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHuancayo                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HUANCAYO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHuanuco                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HUANUCO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHull                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HULL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHurlingham              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HURLINGHAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityHyderabad               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "HYDERABAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIasi                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "IASI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIbague                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "IBAGUE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIca                     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ICA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIlam                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ILAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIlford                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ILFORD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIligan                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ILIGAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIloiloCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ILOILO_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityImperatriz              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "IMPERATRIZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityImus                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "IMUS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIncheon                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "INCHEON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIndaiatuba              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "INDAIATUBA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIndianapolis            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "INDIANAPOLIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIndore                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "INDORE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIpatinga                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "IPATINGA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIpoh                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "IPOH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIquique                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "IQUIQUE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIrvine                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "IRVINE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIsidroCasanova          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ISIDRO_CASANOVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIslamabad               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ISLAMABAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIslington               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ISLINGTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIsmailia                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ISMAILIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIsparta                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ISPARTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIstanbul                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ISTANBUL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityItaborai                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ITABORAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityItabuna                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ITABUNA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityItajai                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ITAJAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityItanhaem                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ITANHAEM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityItapevi                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ITAPEVI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityItaquaquecetuba         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ITAQUAQUECETUBA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityItuzaingo               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ITUZAINGO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIzmir                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "IZMIR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityIztapalapa              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "IZTAPALAPA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJaboataoDosGuararapes   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JABOATAO_DOS GUARARAPES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJacarei                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JACAREI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJackson                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JACKSON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJacksonville            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JACKSONVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJaipur                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JAIPUR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJakarta                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JAKARTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJaraguaDoSul            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JARAGUA_DO SUL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJau                     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JAU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJeddah                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JEDDAH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJember                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JEMBER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJerusalem               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JERUSALEM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJoaoMonlevade           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JOAO_MONLEVADE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJoaoPessoa              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JOAO_PESSOA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJodhpur                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JODHPUR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJohannesburg            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JOHANNESBURG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJohorBahru              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JOHOR_BAHRU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJoinville               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JOINVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJoseCPaz                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JOSE_C PAZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJoseMariaEzeiza         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JOSE_MARIA EZEIZA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJuarez                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JUAREZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJuazeiroDoNorte         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JUAZEIRO_DO NORTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJuizDeFora              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JUIZ_DE FORA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityJundiai                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "JUNDIAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKahramanmaras           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KAHRAMANMARAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKampala                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KAMPALA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKanpur                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KANPUR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKansasCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KANSAS_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKaohsiungCity           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KAOHSIUNG_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKarabuk                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KARABUK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKarachi                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KARACHI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKarlsruhe               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KARLSRUHE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKarnal                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KARNAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKaski                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KASKI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKastamonu               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KASTAMONU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKathmandu               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KATHMANDU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKatowice                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KATOWICE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKatsina                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KATSINA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKaty                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KATY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKaunas                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KAUNAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKayseri                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KAYSERI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKazan                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KAZAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKecskemet               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KECSKEMET"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKediri                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KEDIRI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKenitra                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KENITRA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKharkiv                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KHARKIV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKhmelnytskyi            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KHMELNYTSKYI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKhonKaen                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KHON_KAEN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKielce                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KIELCE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKigali                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KIGALI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKingston                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KINGSTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKirklareli              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KIRKLARELI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKissimmee               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KISSIMMEE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKitchener               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KITCHENER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKlaipeda                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KLAIPEDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKnoxville               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KNOXVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKochi                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KOCHI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKolkata                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KOLKATA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKollam                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KOLLAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKonya                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KONYA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKosekoy                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KOSEKOY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKosice                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KOSICE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKotaKinabalu            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KOTA_KINABALU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKozhikode               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KOZHIKODE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKrakow                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KRAKOW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKrasnodar               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KRASNODAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKryvyiRih               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KRYVYI_RIH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKualaLumpur             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KUALA_LUMPUR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKuching                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KUCHING"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKutahya                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KUTAHYA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKutaisi                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KUTAISI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKuwaitCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KUWAIT_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityKyiv                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "KYIV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLaPaz                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LA_PAZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLaPlata                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LA_PLATA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLaRioja                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LA_RIOJA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLaSerena                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LA_SERENA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLafayette               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAFAYETTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLaferrere               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAFERRERE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLages                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAGES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLagos                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAGOS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLahore                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAHORE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLahug                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAHUG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLakeWorth               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAKE_WORTH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLakeland                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAKELAND"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLancaster               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LANCASTER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLanus                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LANUS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLasPalmasDeGranCanaria  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAS_PALMAS DE GRAN CANARIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLasPinas                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAS_PINAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLasVegas                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAS_VEGAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLausanne                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAUSANNE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLaval                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAVAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLawrenceville           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LAWRENCEVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLeMans                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LE_MANS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLeeds                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LEEDS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLeicester               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LEICESTER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLeipzig                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LEIPZIG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLeon                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LEON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLexington               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LEXINGTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLibreville              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LIBREVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLiege                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LIEGE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLille                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLima                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LIMA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLimassol                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LIMASSOL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLimeira                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LIMEIRA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLincoln                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LINCOLN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLinhares                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LINHARES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLipaCity                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LIPA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLisbon                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LISBON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLiverpool               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LIVERPOOL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLjubljana               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LJUBLJANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLodz                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LODZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLoja                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LOJA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLomasDeZamora           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LOMAS_DE ZAMORA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLome                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LOME"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLondrina                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LONDRINA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLongBeach               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LONG_BEACH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLongueuil               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LONGUEUIL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLouisville              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LOUISVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLuanda                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LUANDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLublin                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LUBLIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLucenaCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LUCENA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLucknow                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LUCKNOW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLudhiana                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LUDHIANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLusaka                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LUSAKA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLuxembourg              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LUXEMBOURG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLuziania                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LUZIANIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLviv                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LVIV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityLyon                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "LYON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMabalacat               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MABALACAT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMacae                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MACAE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMacao                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MACAO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMacapa                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MACAPA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMaceio                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MACEIO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMachala                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MACHALA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMadison                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MADISON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMadrid                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MADRID"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMage                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MAGE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMagelang                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MAGELANG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMagnesiaAdSipylum       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MAGNESIA_AD SIPYLUM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMakassar                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MAKASSAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMakatiCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MAKATI_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMalabon                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MALABON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMalaga                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MALAGA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMalang                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MALANG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMalappuram              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MALAPPURAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMaldonado               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MALDONADO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMale                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MALE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMalmo                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MALMO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityManado                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MANADO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityManagua                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MANAGUA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityManama                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MANAMA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityManaus                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MANAUS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityManchester              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MANCHESTER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMandaluyongCity         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MANDALUYONG_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityManila                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MANILA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityManizales               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MANIZALES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMannheim                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MANNHEIM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMaputo                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MAPUTO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMarDelPlata             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MAR_DEL PLATA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMaraba                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARABA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMaracaibo               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARACAIBO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMaracanau               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARACANAU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMaracay                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARACAY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMardin                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARDIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMaribor                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARIBOR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMarica                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARICA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMarietta                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARIETTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMarikinaCity            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARIKINA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMarilia                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARILIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMaringa                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARINGA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMarrakesh               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARRAKESH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMarseille               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MARSEILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMaua                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MAUA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMazatlan                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MAZATLAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMedan                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MEDAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMedellin                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MEDELLIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMedina                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MEDINA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMeerut                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MEERUT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMeknes                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MEKNES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMelbourne               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MELBOURNE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMemphis                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MEMPHIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMendoza                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MENDOZA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMerida                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MERIDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMerkez                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MERKEZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMerlo                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MERLO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMersin                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MERSIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMesa                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MESA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMexicali                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MEXICALI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMexicoCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MEXICO_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMilan                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MILAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMiltonKeynes            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MILTON_KEYNES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMilwaukee               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MILWAUKEE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMinneapolis             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MINNEAPOLIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMinsk                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MINSK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMiskolc                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MISKOLC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMississauga             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MISSISSAUGA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMogiDasCruzes           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MOGI_DAS CRUZES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMohali                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MOHALI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMonroe                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MONROE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMonteGrande             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MONTE_GRANDE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMontegoBay              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MONTEGO_BAY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMonterrey               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MONTERREY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMontesClaros            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MONTES_CLAROS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMontevideo              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MONTEVIDEO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMontgomery              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MONTGOMERY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMontpellier             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MONTPELLIER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMontreal                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MONTREAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMorelia                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MORELIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMoreno                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MORENO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMoron                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MORON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMossoro                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MOSSORO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMugla                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MUGLA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMultan                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MULTAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMumbai                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MUMBAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMunich                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MUNICH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMurcia                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MURCIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMuscat                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MUSCAT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMuzaffargarh            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MUZAFFARGARH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityMykolayiv               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "MYKOLAYIV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNaaldwijk               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NAALDWIJK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNaga                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NAGA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNagpur                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NAGPUR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNairobi                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NAIROBI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNantes                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NANTES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNaples                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NAPLES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNashville               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NASHVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNassau                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NASSAU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNasugbu                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NASUGBU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNatal                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NATAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNaucalpan               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NAUCALPAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNaviMumbai              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NAVI_MUMBAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNeiva                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NEIVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNeuquen                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NEUQUEN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNevsehir                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NEVSEHIR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNewDelhi                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NEW_DELHI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNewOrleans              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NEW_ORLEANS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNewTaipei               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NEW_TAIPEI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNewark                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NEWARK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNewcastleUponTyne       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NEWCASTLE_UPON TYNE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNhaTrang                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NHA_TRANG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNice                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NICE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNicosia                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NICOSIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNilopolis               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NILOPOLIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNis                     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNiteroi                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NITEROI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNitra                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NITRA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNizhniyNovgorod         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NIZHNIY_NOVGOROD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNogales                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NOGALES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNoida                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NOIDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNorthampton             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NORTHAMPTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNorwich                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NORWICH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNottingham              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NOTTINGHAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNovaFriburgo            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NOVA_FRIBURGO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNovaIguacu              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NOVA_IGUACU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNoviSad                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NOVI_SAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNovoHamburgo            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NOVO_HAMBURGO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNovosibirsk             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NOVOSIBIRSK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityNuremberg               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "NUREMBERG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOakland                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OAKLAND"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOaxacaCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OAXACA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOdesa                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ODESA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOklahomaCity            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OKLAHOMA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOlinda                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OLINDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOlomouc                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OLOMOUC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOlongapoCity            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OLONGAPO_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOlsztyn                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OLSZTYN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOmaha                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OMAHA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOradea                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ORADEA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOran                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ORAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOrdu                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ORDU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOrlando                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ORLANDO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOsasco                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OSASCO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOslo                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OSLO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOsmaniye                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OSMANIYE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOstrava                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OSTRAVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOttawa                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OTTAWA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOujda                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OUJDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityOurinhos                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "OURINHOS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPachuca                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PACHUCA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPadova                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PADOVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPalakkad                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PALAKKAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPalembang               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PALEMBANG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPalermo                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PALERMO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPalhoca                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PALHOCA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPalma                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PALMA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPalmas                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PALMAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPanamaCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PANAMA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityParamaribo              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PARAMARIBO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityParana                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PARANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityParanagua               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PARANAGUA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityParanaqueCity           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PARANAQUE_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityParauapebas             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PARAUAPEBAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityParis                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PARIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityParnaiba                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PARNAIBA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityParnamirim              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PARNAMIRIM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPassoFundo              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PASSO_FUNDO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPasto                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PASTO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPatan                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PATAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPatna                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PATNA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPatosDeMinas            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PATOS_DE MINAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPaulista                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PAULISTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPecs                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PECS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPekanbaru               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PEKANBARU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPelotas                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PELOTAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPeoria                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PEORIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPereira                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PEREIRA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPerm                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PERM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPerth                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PERTH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPescara                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PESCARA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPeshawar                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PESHAWAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPetahTikva              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PETAH_TIKVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPetalingJaya            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PETALING_JAYA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPetrolina               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PETROLINA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPetropolis              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PETROPOLIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPhiladelphia            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PHILADELPHIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPhnomPenh               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PHNOM_PENH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPhoenix                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PHOENIX"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPilar                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PILAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPindamonhangaba         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PINDAMONHANGABA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPiracicaba              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PIRACICABA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPitesti                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PITESTI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPittsburgh              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PITTSBURGH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPiura                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PIURA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPlano                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PLANO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPloiesti                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PLOIESTI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPlovdiv                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PLOVDIV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPlymouth                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PLYMOUTH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPocosDeCaldas           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "POCOS_DE CALDAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPodgorica               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PODGORICA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPoltava                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "POLTAVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPontaGrossa             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PONTA_GROSSA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPontianak               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PONTIANAK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPopayan                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "POPAYAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortAuPrince            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORT_AU PRINCE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortElizabeth           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORT_ELIZABETH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortHarcourt            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORT_HARCOURT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortLouis               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORT_LOUIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortMontt               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORT_MONTT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortOfSpain             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORT_OF SPAIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortSaid                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORT_SAID"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortland                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORTLAND"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPorto                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORTO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortoAlegre             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORTO_ALEGRE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortoSeguro             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORTO_SEGURO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortoVelho              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORTO_VELHO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPortoviejo              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PORTOVIEJO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPosadas                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "POSADAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPousoAlegre             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "POUSO_ALEGRE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPoznan                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "POZNAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPrague                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PRAGUE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPraiaGrande             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PRAIA_GRANDE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPresidentePrudente      CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PRESIDENTE_PRUDENTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPretoria                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PRETORIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPristina                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PRISTINA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityProvidence              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PROVIDENCE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPucallpa                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PUCALLPA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPuchongBatuDuaBelas     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PUCHONG_BATU DUA BELAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPueblaCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PUEBLA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityPune                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "PUNE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityQuebec                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "QUEBEC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityQueens                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "QUEENS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityQueimados               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "QUEIMADOS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityQueretaroCity           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "QUERETARO_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityQuezonCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "QUEZON_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityQuilmes                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "QUILMES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityQuito                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "QUITO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRabat                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RABAT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRaipur                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RAIPUR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRajkot                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RAJKOT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRajshahi                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RAJSHAHI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRaleigh                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RALEIGH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRamatGan                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RAMAT_GAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRancagua                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RANCAGUA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRanchi                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RANCHI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRasAlKhaimah            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RAS_AL KHAIMAH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRawalpindi              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RAWALPINDI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityReading                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "READING"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRecife                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RECIFE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRegina                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "REGINA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRennes                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RENNES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityReno                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RENO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityResistencia             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RESISTENCIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityReykjavik               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "REYKJAVIK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityReynosa                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "REYNOSA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRibeiraoDasNeves        CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIBEIRAO_DAS NEVES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRibeiraoPreto           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIBEIRAO_PRETO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRichmond                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RICHMOND"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRiga                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIGA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRioBranco               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIO_BRANCO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRioClaro                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIO_CLARO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRioCuarto               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIO_CUARTO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRioDeJaneiro            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIO_DE JANEIRO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRioDoSul                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIO_DO SUL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRioGallegos             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIO_GALLEGOS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRioGrande               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIO_GRANDE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRishonLetsiyyon         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RISHON_LETSIYYON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRiverside               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIVERSIDE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRiyadh                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIYADH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRize                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RIZE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRochester               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ROCHESTER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRome                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ROME"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRondonopolis            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RONDONOPOLIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRosario                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ROSARIO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRoseau                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ROSEAU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRostovOnDon             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ROSTOV_ON DON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRotterdam               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ROTTERDAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRouen                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ROUEN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRousse                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ROUSSE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityRzeszow                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "RZESZOW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySacramento              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SACRAMENTO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySagar                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAGAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaintPaul               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAINT_PAUL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySale                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SALE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaltLakeCity            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SALT_LAKE CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySalta                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SALTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaltillo                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SALTILLO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySalvador                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SALVADOR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySamara                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAMARA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySamarinda               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAMARINDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySamarkand               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAMARKAND"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySamsun                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAMSUN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanAntonio              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_ANTONIO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanDiego                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_DIEGO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanFernando             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_FERNANDO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanFrancisco            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_FRANCISCO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanJose                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_JOSE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanJoseDelMonte         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_JOSE DEL MONTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanJuan                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_JUAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanJusto                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_JUSTO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanLuis                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_LUIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanLuisPotosiCity       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_LUIS POTOSI CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanMiguel               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_MIGUEL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanMiguelDeTucuman      CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_MIGUEL DE TUCUMAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanPabloCity            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_PABLO CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanPedro                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_PEDRO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanPedroSula            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_PEDRO SULA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanSalvador             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_SALVADOR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanSalvadorDeJujuy      CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAN_SALVADOR DE JUJUY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanaa                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANAA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySanliurfa               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANLIURFA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantaCruz               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTA_CRUZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantaCruzDeTenerife     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTA_CRUZ DE TENERIFE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantaCruzDoSul          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTA_CRUZ DO SUL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantaFe                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTA_FE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantaLuzia              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTA_LUZIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantaMaria              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTA_MARIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantaMarta              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTA_MARTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantaRosa               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTA_ROSA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantarem                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTAREM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantiago                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTIAGO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantiagoDeCali          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTIAGO_DE CALI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantiagoDeLosCaballeros CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTIAGO_DE LOS CABALLEROS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantoAndre              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTO_ANDRE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantoDomingo            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTO_DOMINGO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantoDomingoEste        CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTO_DOMINGO ESTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySantos                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SANTOS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoBernardoDoCampo      CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_BERNARDO DO CAMPO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoCarlos               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_CARLOS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoGoncalo              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_GONCALO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoJoaoDeMeriti         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_JOAO DE MERITI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoJose                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_JOSE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoJoseDoRioPreto       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_JOSE DO RIO PRETO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoJoseDosCampos        CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_JOSE DOS CAMPOS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoJoseDosPinhais       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_JOSE DOS PINHAIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoLeopoldo             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_LEOPOLDO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoLuis                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_LUIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoPaulo                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_PAULO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaoVicente              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SAO_VICENTE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySarajevo                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SARAJEVO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySaskatoon               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SASKATOON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityScarborough             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SCARBOROUGH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySeattle                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SEATTLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySemarang                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SEMARANG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySeoGu                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SEO_GU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySeongnamSi              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SEONGNAM_SI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySeoul                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SEOUL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySerra                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SERRA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySeteLagoas              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SETE_LAGOAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySetif                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SETIF"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySetubal                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SETUBAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySeville                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SEVILLE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySfax                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SFAX"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityShahAlam                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SHAH_ALAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityShanghai                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SHANGHAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySharjah                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SHARJAH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySheffield               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SHEFFIELD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityShenzhen                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SHENZHEN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityShimla                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SHIMLA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySiauliai                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SIAULIAI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySibiu                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SIBIU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySidoarjo                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SIDOARJO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySikar                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SIKAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySilverSpring            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SILVER_SPRING"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySinop                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SINOP"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySivas                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SIVAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySkikda                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SKIKDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySkopje                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SKOPJE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySlough                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SLOUGH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySobral                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SOBRAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySofia                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SOFIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySorocaba                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SOROCABA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySousse                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SOUSSE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySouthTangerang          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SOUTH_TANGERANG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySouthampton             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SOUTHAMPTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySouthwark               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SOUTHWARK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySplit                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SPLIT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySpokane                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SPOKANE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySpring                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SPRING"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySpringfield             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SPRINGFIELD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityStLouis                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ST_LOUIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityStPetersburg            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ST_PETERSBURG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityStaraZagora             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "STARA_ZAGORA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityStatenIsland            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "STATEN_ISLAND"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityStockholm               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "STOCKHOLM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityStockton                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "STOCKTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityStokeOnTrent            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "STOKE_ON TRENT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityStrasbourg              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "STRASBOURG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityStuttgart               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "STUTTGART"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySumare                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SUMARE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySurabaya                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SURABAYA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySurakarta               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SURAKARTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySurat                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SURAT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySurrey                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SURREY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySuwon                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SUWON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySuzano                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SUZANO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySydney                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SYDNEY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySzczecin                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SZCZECIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySzeged                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SZEGED"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCitySzekesfehervar          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "SZEKESFEHERVAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTaboaoDaSerra           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TABOAO_DA SERRA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTacna                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TACNA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTacoma                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TACOMA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTaguig                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TAGUIG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTaichung                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TAICHUNG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTainanCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TAINAN_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTaipei                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TAIPEI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTalavera                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TALAVERA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTalca                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TALCA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTallahassee             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TALLAHASSEE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTallinn                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TALLINN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTampa                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TAMPA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTampere                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TAMPERE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTampico                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TAMPICO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTangerang               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TANGERANG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTangier                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TANGIER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTanta                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TANTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTanza                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TANZA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTaoyuanDistrict         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TAOYUAN_DISTRICT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTappahannock            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TAPPAHANNOCK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTarlacCity              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TARLAC_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTashkent                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TASHKENT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTasikmalaya             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TASIKMALAYA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTatui                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TATUI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTaubate                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TAUBATE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTbilisi                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TBILISI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTegucigalpa             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TEGUCIGALPA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTehran                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TEHRAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTeixeiraDeFreitas       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TEIXEIRA_DE FREITAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTekirdag                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TEKIRDAG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTelAviv                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TEL_AVIV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTemuco                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TEMUCO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTepic                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TEPIC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTeresina                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TERESINA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTernopil                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TERNOPIL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTerrassa                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TERRASSA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTetouan                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TETOUAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityThane                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "THANE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTheBronx                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "THE_BRONX"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTheHague                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "THE_HAGUE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityThessaloniki            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "THESSALONIKI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityThiruvananthapuram      CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "THIRUVANANTHAPURAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityThrissur                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "THRISSUR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTijuana                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TIJUANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTimisoara               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TIMISOARA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTirana                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TIRANA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTlalnepantla            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TLALNEPANTLA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTlaxcalaCity            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TLAXCALA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTlemcen                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TLEMCEN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTokatProvince           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TOKAT_PROVINCE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTokyo                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TOKYO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityToluca                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TOLUCA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityToronto                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TORONTO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTorreon                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TORREON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityToulouse                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TOULOUSE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTrabzon                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TRABZON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTrujillo                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TRUJILLO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTubarao                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TUBARAO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTucson                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TUCSON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTuguegaraoCity          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TUGUEGARAO_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTulsa                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TULSA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTunis                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TUNIS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTunja                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TUNJA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTurin                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TURIN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTuxtlaGutierrez         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TUXTLA_GUTIERREZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityTuzla                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "TUZLA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityUberaba                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "UBERABA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityUberlandia              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "UBERLANDIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityUfa                     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "UFA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityUlanBator               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ULAN_BATOR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityUmeda                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "UMEDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityUrdaneta                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "URDANETA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityUsak                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "USAK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVadodara                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VADODARA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityValencia                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VALENCIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityValinhos                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VALINHOS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityValladolid              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VALLADOLID"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityValledupar              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VALLEDUPAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityValparaiso              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VALPARAISO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityValparaisoDeGoias       CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VALPARAISO_DE GOIAS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVan                     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVancouver               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VANCOUVER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVaranasi                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VARANASI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVarginha                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VARGINHA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVarna                   CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VARNA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVarzeaPaulista          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VARZEA_PAULISTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVenustianoCarranza      CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VENUSTIANO_CARRANZA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVeracruz                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VERACRUZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVerona                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VERONA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityViamao                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VIAMAO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVictoria                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VICTORIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVienna                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VIENNA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVientiane               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VIENTIANE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVigo                    CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VIGO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVijayawada              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VIJAYAWADA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVilaNovaDeGaia          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VILA_NOVA DE GAIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVilaVelha               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VILA_VELHA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVillaBallester          CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VILLA_BALLESTER"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVillavicencio           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VILLAVICENCIO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVilnius                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VILNIUS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVinaDelMar              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VINA_DEL MAR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVinnytsia               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VINNYTSIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVirginiaBeach           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VIRGINIA_BEACH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVisakhapatnam           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VISAKHAPATNAM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVitoria                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VITORIA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVitoriaDaConquista      CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VITORIA_DA CONQUISTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVitoriaDeSantoAntao     CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VITORIA_DE SANTO ANTAO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVoltaRedonda            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VOLTA_REDONDA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityVoronezh                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "VORONEZH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWarsaw                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WARSAW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWashington              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WASHINGTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWellington              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WELLINGTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWestPalmBeach           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WEST_PALM BEACH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWichita                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WICHITA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWillemstad              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WILLEMSTAD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWilmington              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WILMINGTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWindhoek                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WINDHOEK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWindsor                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WINDSOR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWinnipeg                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WINNIPEG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWolverhampton           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WOLVERHAMPTON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWoodbridge              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WOODBRIDGE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWroclaw                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WROCLAW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityWuppertal               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "WUPPERTAL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityXalapa                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "XALAPA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityYalova                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "YALOVA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityYangon                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "YANGON"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityYekaterinburg           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "YEKATERINBURG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityYerevan                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "YEREVAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityYogyakarta              CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "YOGYAKARTA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityYokohama                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "YOKOHAMA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityYonginSi                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "YONGIN_SI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZabrze                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZABRZE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZagazig                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZAGAZIG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZagreb                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZAGREB"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZamboangaCity           CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZAMBOANGA_CITY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZapopan                 CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZAPOPAN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZaporizhzhya            CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZAPORIZHZHYA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZaragoza                CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZARAGOZA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZhongliDistrict         CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZHONGLI_DISTRICT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZielonaGora             CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZIELONA_GORA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZonguldak               CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZONGULDAK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCityZurich                  CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity = "ZURICH"
)

type CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry string

const (
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryUs CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "US"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryCa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "CA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMx CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MX"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryGB CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "GB"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryDe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "DE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryFr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "FR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryIt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "IT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryEs CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "ES"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryPl CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "PL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryNl CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "NL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryNo CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "NO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryDk CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "DK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryFi CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "FI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryCh CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "CH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryAt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "AT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryIe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "IE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryPt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "PT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryGr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "GR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryCz CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "CZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryHu CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "HU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryRo CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "RO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBg CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySk CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySi CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryHr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "HR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryEe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "EE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryLv CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "LV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryLt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "LT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryLu CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "LU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryCy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "CY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryIs CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "IS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryLi CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "LI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMc CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySm CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryVa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "VA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryJp CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "JP"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryKr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "KR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryCn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "CN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryHk CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "HK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryTw CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "TW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySg CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryAu CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "AU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryNz CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "NZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryIn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "IN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryTh CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "TH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryPh CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "PH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryID CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "ID"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryVn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "VN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryAf CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "AF"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBd CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryKh CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "KH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryLa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "LA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryLk CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "LK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMm CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryNp CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "NP"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryPk CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "PK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryFj CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "FJ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryPg CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "PG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryAe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "AE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryIl CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "IL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryTr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "TR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryIr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "IR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryIq CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "IQ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryJo CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "JO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryKw CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "KW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryLb CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "LB"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryOm CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "OM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryQa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "QA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBh CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryYe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "YE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryZa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "ZA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryEg CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "EG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryNg CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "NG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryKe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "KE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryDz CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "DZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryAo CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "AO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBw CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryEt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "ET"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryGh CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "GH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryCi CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "CI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryLy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "LY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMz CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryRw CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "RW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryTn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "TN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryUg CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "UG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryZm CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "ZM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryZw CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "ZW"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryTz CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "TZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMu CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySc CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryAr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "AR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryCl CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "CL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryCo CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "CO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryPe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "PE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryVe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "VE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryEc CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "EC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryUy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "UY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryPy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "PY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBo CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryCr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "CR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryCu CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "CU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryDo CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "DO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryGt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "GT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryHn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "HN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryJm CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "JM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryNi CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "NI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryPa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "PA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySv CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryTt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "TT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBb CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BB"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBz CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryGy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "GY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountrySr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "SR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryRu CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "RU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryUa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "UA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryKz CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "KZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryUz CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "UZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryAz CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "AZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryGe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "GE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryAm CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "AM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMd CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMk CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryAl CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "AL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryBa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "BA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryRs CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "RS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "ME"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryXk CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "XK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryMn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "MN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryKg CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "KG"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryTj CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "TJ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountryTm CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry = "TM"
)

type CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState string

const (
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateAl CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "AL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateAk CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "AK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateAz CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "AZ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateAr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "AR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateCa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "CA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateCo CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "CO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateCt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "CT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateDe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "DE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateFl CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "FL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateGa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "GA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateHi CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "HI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateID CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "ID"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateIl CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "IL"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateIn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "IN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateIa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "IA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateKs CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "KS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateKy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "KY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateLa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "LA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateMe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "ME"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateMd CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "MD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateMa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "MA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateMi CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "MI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateMn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "MN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateMs CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "MS"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateMo CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "MO"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateMt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "MT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateNe CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "NE"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateNv CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "NV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateNh CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "NH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateNj CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "NJ"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateNm CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "NM"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateNy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "NY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateNc CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "NC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateNd CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "ND"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateOh CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "OH"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateOk CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "OK"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateOr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "OR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStatePa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "PA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateRi CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "RI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateSc CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "SC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateSd CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "SD"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateTn CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "TN"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateTx CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "TX"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateUt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "UT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateVt CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "VT"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateVa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "VA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateWa CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "WA"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateWv CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "WV"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateWi CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "WI"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateWy CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "WY"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateDc CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "DC"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStatePr CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "PR"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateGu CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "GU"
	CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationStateVi CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState = "VI"
)

type SessionCostResponseUsageBrowserUnit string

const (
	SessionCostResponseUsageBrowserUnitMinute SessionCostResponseUsageBrowserUnit = "minute"
	SessionCostResponseUsageBrowserUnitSecond SessionCostResponseUsageBrowserUnit = "second"
)

// ErrorResponse An error response from the API
type ErrorResponse struct {
	Context    []ErrorResponseContext `json:"context"`
	Error      string                 `json:"error"`
	LinkToDocs string                 `json:"linkToDocs"`
	Message    string                 `json:"message" api:"required"`
	JSON       errorResponseJSON      `json:"-"`
}

// errorResponseJSON contains the JSON metadata for the struct [ErrorResponse]
type errorResponseJSON struct {
	Context     apijson.Field
	Error       apijson.Field
	LinkToDocs  apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ErrorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r errorResponseJSON) RawJSON() string {
	return r.raw
}

type File struct {
	// Timestamp when the file was created
	LastModified time.Time `json:"lastModified" api:"required"`
	// Path to the file in the storage system
	Path string `json:"path" api:"required"`
	// Size of the file in bytes
	Size int64    `json:"size" api:"required"`
	JSON fileJSON `json:"-"`
}

// fileJSON contains the JSON metadata for the struct [File]
type fileJSON struct {
	LastModified apijson.Field
	Path         apijson.Field
	Size         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *File) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileJSON) RawJSON() string {
	return r.raw
}

type Fileslist struct {
	// Array of files for the current page
	Data []FileslistData `json:"data" api:"required"`
	JSON fileslistJSON   `json:"-"`
}

// fileslistJSON contains the JSON metadata for the struct [Fileslist]
type fileslistJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Fileslist) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileslistJSON) RawJSON() string {
	return r.raw
}

type ClientPdfParams struct {
	// Delay before generating the PDF (in milliseconds)
	Delay param.Field[float64] `json:"delay"`
	// Project to execute the PDF generation in.
	ProjectID param.Field[string] `json:"projectId"`
	// The desired region for the action to be performed in
	Region param.Field[string] `json:"region"`
	// URL of the webpage to convert to PDF
	URL param.Field[string] `json:"url" api:"required"`
	// Use a Steel-provided residential proxy for generating the PDF
	UseProxy param.Field[bool] `json:"useProxy"`
}

func (r ClientPdfParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PdfResponse struct {
	// URL where the PDF is hosted
	URL  string          `json:"url" api:"required"`
	JSON pdfResponseJSON `json:"-"`
}

// pdfResponseJSON contains the JSON metadata for the struct [PdfResponse]
type pdfResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PdfResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pdfResponseJSON) RawJSON() string {
	return r.raw
}

type ClientScrapeParams struct {
	// Delay before scraping (in milliseconds)
	Delay param.Field[float64] `json:"delay"`
	// Desired format(s) for the scraped content. Default is `html`.
	Format param.Field[[]ScrapeRequestFormatItem] `json:"format"`
	// Include a PDF in the response
	Pdf param.Field[bool] `json:"pdf"`
	// Project to execute the scrape in.
	ProjectID param.Field[string] `json:"projectId"`
	// The desired region for the action to be performed in
	Region param.Field[string] `json:"region"`
	// Include a screenshot in the response
	Screenshot param.Field[bool] `json:"screenshot"`
	// URL of the webpage to scrape
	URL param.Field[string] `json:"url" api:"required"`
	// Use a Steel-provided residential proxy for the scrape
	UseProxy param.Field[bool] `json:"useProxy"`
}

func (r ClientScrapeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ScrapeResponse Response from a successful scrape request
type ScrapeResponse struct {
	Content    ScrapeResponseContent    `json:"content" api:"required"`
	Links      []ScrapeResponseLink     `json:"links" api:"required"`
	Metadata   ScrapeResponseMetadata   `json:"metadata" api:"required"`
	Pdf        ScrapeResponsePdf        `json:"pdf"`
	Screenshot ScrapeResponseScreenshot `json:"screenshot"`
	JSON       scrapeResponseJSON       `json:"-"`
}

// scrapeResponseJSON contains the JSON metadata for the struct [ScrapeResponse]
type scrapeResponseJSON struct {
	Content     apijson.Field
	Links       apijson.Field
	Metadata    apijson.Field
	Pdf         apijson.Field
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeResponseJSON) RawJSON() string {
	return r.raw
}

type ClientScreenshotParams struct {
	// Delay before capturing the screenshot (in milliseconds)
	Delay param.Field[float64] `json:"delay"`
	// Capture the full page screenshot. Default is `false`.
	FullPage param.Field[bool] `json:"fullPage"`
	// Project to execute the screenshot in.
	ProjectID param.Field[string] `json:"projectId"`
	// The desired region for the action to be performed in
	Region param.Field[string] `json:"region"`
	// URL of the webpage to capture
	URL param.Field[string] `json:"url" api:"required"`
	// Use a Steel-provided residential proxy for capturing the screenshot
	UseProxy param.Field[bool] `json:"useProxy"`
}

func (r ClientScreenshotParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotResponse struct {
	// URL where the screenshot is hosted
	URL  string                 `json:"url" api:"required"`
	JSON screenshotResponseJSON `json:"-"`
}

// screenshotResponseJSON contains the JSON metadata for the struct [ScreenshotResponse]
type screenshotResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScreenshotResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r screenshotResponseJSON) RawJSON() string {
	return r.raw
}

type ErrorResponseContext struct {
	Keyword string                   `json:"keyword" api:"required"`
	Message string                   `json:"message" api:"required"`
	Params  interface{}              `json:"params" api:"required"`
	JSON    errorResponseContextJSON `json:"-"`
}

// errorResponseContextJSON contains the JSON metadata for the struct [ErrorResponseContext]
type errorResponseContextJSON struct {
	Keyword     apijson.Field
	Message     apijson.Field
	Params      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ErrorResponseContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r errorResponseContextJSON) RawJSON() string {
	return r.raw
}

// FileslistData Array of files for the current page
type FileslistData struct {
	// Timestamp when the file was created
	LastModified time.Time `json:"lastModified" api:"required"`
	// Path to the file in the storage system
	Path string `json:"path" api:"required"`
	// Size of the file in bytes
	Size int64             `json:"size" api:"required"`
	JSON fileslistDataJSON `json:"-"`
}

// fileslistDataJSON contains the JSON metadata for the struct [FileslistData]
type fileslistDataJSON struct {
	LastModified apijson.Field
	Path         apijson.Field
	Size         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FileslistData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileslistDataJSON) RawJSON() string {
	return r.raw
}

type ScrapeResponseContent struct {
	// Cleaned HTML content of the webpage
	CleanedHTML string `json:"cleaned_html"`
	// Raw HTML content of the webpage
	HTML string `json:"html"`
	// Webpage content converted to Markdown
	Markdown string `json:"markdown"`
	// Webpage content in Readability format
	Readability map[string]interface{}    `json:"readability"`
	JSON        scrapeResponseContentJSON `json:"-"`
}

// scrapeResponseContentJSON contains the JSON metadata for the struct [ScrapeResponseContent]
type scrapeResponseContentJSON struct {
	CleanedHTML apijson.Field
	HTML        apijson.Field
	Markdown    apijson.Field
	Readability apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeResponseContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeResponseContentJSON) RawJSON() string {
	return r.raw
}

type ScrapeResponseLink struct {
	// Text content of the link
	Text string `json:"text" api:"required"`
	// URL of the link
	URL  string                 `json:"url" api:"required"`
	JSON scrapeResponseLinkJSON `json:"-"`
}

// scrapeResponseLinkJSON contains the JSON metadata for the struct [ScrapeResponseLink]
type scrapeResponseLinkJSON struct {
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeResponseLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeResponseLinkJSON) RawJSON() string {
	return r.raw
}

type ScrapeResponseMetadata struct {
	// Author of the article content
	ArticleAuthor string `json:"articleAuthor"`
	// Author of the webpage content
	Author string `json:"author"`
	// Canonical URL of the webpage
	Canonical string `json:"canonical"`
	// Description of the webpage
	Description string `json:"description"`
	// Favicon URL of the website
	Favicon string `json:"favicon"`
	// JSON-LD structured data from the webpage
	JSONLd interface{} `json:"jsonLd"`
	// Keywords associated with the webpage
	Keywords string `json:"keywords"`
	// Detected language of the webpage
	Language string `json:"language"`
	// Last modification time of the content
	ModifiedTime string `json:"modifiedTime"`
	// Open Graph description
	OgDescription string `json:"ogDescription"`
	// Open Graph image URL
	OgImage string `json:"ogImage"`
	// Open Graph site name
	OgSiteName string `json:"ogSiteName"`
	// Open Graph title
	OgTitle string `json:"ogTitle"`
	// Open Graph URL
	OgURL string `json:"ogUrl"`
	// Publication time of the content
	PublishedTime string `json:"publishedTime"`
	// HTTP status code of the response
	StatusCode int64 `json:"statusCode" api:"required"`
	// Timestamp when the scrape was performed
	Timestamp time.Time `json:"timestamp"`
	// Title of the webpage
	Title string `json:"title"`
	// Source URL of the scraped page
	URLSource string                     `json:"urlSource"`
	JSON      scrapeResponseMetadataJSON `json:"-"`
}

// scrapeResponseMetadataJSON contains the JSON metadata for the struct [ScrapeResponseMetadata]
type scrapeResponseMetadataJSON struct {
	ArticleAuthor apijson.Field
	Author        apijson.Field
	Canonical     apijson.Field
	Description   apijson.Field
	Favicon       apijson.Field
	JSONLd        apijson.Field
	Keywords      apijson.Field
	Language      apijson.Field
	ModifiedTime  apijson.Field
	OgDescription apijson.Field
	OgImage       apijson.Field
	OgSiteName    apijson.Field
	OgTitle       apijson.Field
	OgURL         apijson.Field
	PublishedTime apijson.Field
	StatusCode    apijson.Field
	Timestamp     apijson.Field
	Title         apijson.Field
	URLSource     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScrapeResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type ScrapeResponsePdf struct {
	// URL of the generated PDF
	URL  string                `json:"url" api:"required"`
	JSON scrapeResponsePdfJSON `json:"-"`
}

// scrapeResponsePdfJSON contains the JSON metadata for the struct [ScrapeResponsePdf]
type scrapeResponsePdfJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeResponsePdf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeResponsePdfJSON) RawJSON() string {
	return r.raw
}

type ScrapeResponseScreenshot struct {
	// URL of the screenshot image
	URL  string                       `json:"url" api:"required"`
	JSON scrapeResponseScreenshotJSON `json:"-"`
}

// scrapeResponseScreenshotJSON contains the JSON metadata for the struct [ScrapeResponseScreenshot]
type scrapeResponseScreenshotJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeResponseScreenshot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeResponseScreenshotJSON) RawJSON() string {
	return r.raw
}

type CaptchaStatusResponse []CaptchaStatusResponseItem

// SessionEventsResponse Events for a browser session
type SessionEventsResponse []interface{}

type SessionReleaseAllParams map[string]interface{}

type SessionReleaseParams map[string]interface{}

// SessionCreateParamsOptimizeBandwidth Enable bandwidth optimizations. Passing true enables all flags (except hosts/patterns). Object allows granular control.
type SessionCreateParamsOptimizeBandwidth = SessionCreateParamsOptimizeBandwidth2

type SessionCreateParamsUseProxy = SessionCreateParamsUseProxy2

type SessionComputerParams struct {
	Action                                 string                                  `json:"action"`
	ComputerActionRequestMoveMouse         *ComputerActionRequestMoveMouse         `json:"-"`
	ComputerActionRequestClickMouse        *ComputerActionRequestClickMouse        `json:"-"`
	ComputerActionRequestDragMouse         *ComputerActionRequestDragMouse         `json:"-"`
	ComputerActionRequestScroll            *ComputerActionRequestScroll            `json:"-"`
	ComputerActionRequestPressKey          *ComputerActionRequestPressKey          `json:"-"`
	ComputerActionRequestTypeText          *ComputerActionRequestTypeText          `json:"-"`
	ComputerActionRequestWait              *ComputerActionRequestWait              `json:"-"`
	ComputerActionRequestTakeScreenshot    *ComputerActionRequestTakeScreenshot    `json:"-"`
	ComputerActionRequestGetCursorPosition *ComputerActionRequestGetCursorPosition `json:"-"`
}

func (u SessionComputerParams) MarshalJSON() ([]byte, error) {
	switch u.Action {
	case "move_mouse":
		return json.Marshal(u.ComputerActionRequestMoveMouse)
	case "click_mouse":
		return json.Marshal(u.ComputerActionRequestClickMouse)
	case "drag_mouse":
		return json.Marshal(u.ComputerActionRequestDragMouse)
	case "scroll":
		return json.Marshal(u.ComputerActionRequestScroll)
	case "press_key":
		return json.Marshal(u.ComputerActionRequestPressKey)
	case "type_text":
		return json.Marshal(u.ComputerActionRequestTypeText)
	case "wait":
		return json.Marshal(u.ComputerActionRequestWait)
	case "take_screenshot":
		return json.Marshal(u.ComputerActionRequestTakeScreenshot)
	case "get_cursor_position":
		return json.Marshal(u.ComputerActionRequestGetCursorPosition)
	}
	return []byte("null"), nil
}

func NewSessionComputerParamsMoveMouse(v ComputerActionRequestMoveMouse) SessionComputerParams {
	return SessionComputerParams{Action: "move_mouse", ComputerActionRequestMoveMouse: &v}
}

func NewSessionComputerParamsClickMouse(v ComputerActionRequestClickMouse) SessionComputerParams {
	return SessionComputerParams{Action: "click_mouse", ComputerActionRequestClickMouse: &v}
}

func NewSessionComputerParamsDragMouse(v ComputerActionRequestDragMouse) SessionComputerParams {
	return SessionComputerParams{Action: "drag_mouse", ComputerActionRequestDragMouse: &v}
}

func NewSessionComputerParamsScroll(v ComputerActionRequestScroll) SessionComputerParams {
	return SessionComputerParams{Action: "scroll", ComputerActionRequestScroll: &v}
}

func NewSessionComputerParamsPressKey(v ComputerActionRequestPressKey) SessionComputerParams {
	return SessionComputerParams{Action: "press_key", ComputerActionRequestPressKey: &v}
}

func NewSessionComputerParamsTypeText(v ComputerActionRequestTypeText) SessionComputerParams {
	return SessionComputerParams{Action: "type_text", ComputerActionRequestTypeText: &v}
}

func NewSessionComputerParamsWait(v ComputerActionRequestWait) SessionComputerParams {
	return SessionComputerParams{Action: "wait", ComputerActionRequestWait: &v}
}

func NewSessionComputerParamsTakeScreenshot(v ComputerActionRequestTakeScreenshot) SessionComputerParams {
	return SessionComputerParams{Action: "take_screenshot", ComputerActionRequestTakeScreenshot: &v}
}

func NewSessionComputerParamsGetCursorPosition(v ComputerActionRequestGetCursorPosition) SessionComputerParams {
	return SessionComputerParams{Action: "get_cursor_position", ComputerActionRequestGetCursorPosition: &v}
}

type SessionCreateParamsOptimizeBandwidth2 struct {
	OfBool                                             *bool                                             `json:"-"`
	OfSessionCreateParamsOptimizeBandwidthUnionMember1 *SessionCreateParamsOptimizeBandwidthUnionMember1 `json:"-"`
}

func (u SessionCreateParamsOptimizeBandwidth2) MarshalJSON() ([]byte, error) {
	if u.OfBool != nil {
		return json.Marshal(u.OfBool)
	}
	if u.OfSessionCreateParamsOptimizeBandwidthUnionMember1 != nil {
		return json.Marshal(u.OfSessionCreateParamsOptimizeBandwidthUnionMember1)
	}
	return []byte("null"), nil
}

type SessionCreateParamsUseProxy2 struct {
	OfBool                                   *bool                                   `json:"-"`
	OfSessionCreateParamsUseProxyGeolocation *SessionCreateParamsUseProxyGeolocation `json:"-"`
	OfSessionCreateParamsUseProxyServer      *SessionCreateParamsUseProxyServer      `json:"-"`
	OfAny                                    *interface{}                            `json:"-"`
}

func (u SessionCreateParamsUseProxy2) MarshalJSON() ([]byte, error) {
	if u.OfBool != nil {
		return json.Marshal(u.OfBool)
	}
	if u.OfSessionCreateParamsUseProxyGeolocation != nil {
		return json.Marshal(u.OfSessionCreateParamsUseProxyGeolocation)
	}
	if u.OfSessionCreateParamsUseProxyServer != nil {
		return json.Marshal(u.OfSessionCreateParamsUseProxyServer)
	}
	if u.OfAny != nil {
		return json.Marshal(u.OfAny)
	}
	return []byte("null"), nil
}
