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

type ComputerActionRequestMoveMouseAction string

const (
	ComputerActionRequestMoveMouseActionMoveMouse ComputerActionRequestMoveMouseAction = "move_mouse"
)

type ComputerActionRequestClickMouseAction string

const (
	ComputerActionRequestClickMouseActionClickMouse ComputerActionRequestClickMouseAction = "click_mouse"
)

type ComputerActionRequestClickMouseButton string

const (
	ComputerActionRequestClickMouseButtonLeft    ComputerActionRequestClickMouseButton = "left"
	ComputerActionRequestClickMouseButtonRight   ComputerActionRequestClickMouseButton = "right"
	ComputerActionRequestClickMouseButtonMiddle  ComputerActionRequestClickMouseButton = "middle"
	ComputerActionRequestClickMouseButtonBack    ComputerActionRequestClickMouseButton = "back"
	ComputerActionRequestClickMouseButtonForward ComputerActionRequestClickMouseButton = "forward"
)

type ComputerActionRequestClickMouseClickType string

const (
	ComputerActionRequestClickMouseClickTypeDown  ComputerActionRequestClickMouseClickType = "down"
	ComputerActionRequestClickMouseClickTypeUp    ComputerActionRequestClickMouseClickType = "up"
	ComputerActionRequestClickMouseClickTypeClick ComputerActionRequestClickMouseClickType = "click"
)

type ComputerActionRequestDragMouseAction string

const (
	ComputerActionRequestDragMouseActionDragMouse ComputerActionRequestDragMouseAction = "drag_mouse"
)

type ComputerActionRequestScrollAction string

const (
	ComputerActionRequestScrollActionScroll ComputerActionRequestScrollAction = "scroll"
)

type ComputerActionRequestPressKeyAction string

const (
	ComputerActionRequestPressKeyActionPressKey ComputerActionRequestPressKeyAction = "press_key"
)

type ComputerActionRequestTypeTextAction string

const (
	ComputerActionRequestTypeTextActionTypeText ComputerActionRequestTypeTextAction = "type_text"
)

type ComputerActionRequestWaitAction string

const (
	ComputerActionRequestWaitActionWait ComputerActionRequestWaitAction = "wait"
)

type ComputerActionRequestTakeScreenshotAction string

const (
	ComputerActionRequestTakeScreenshotActionTakeScreenshot ComputerActionRequestTakeScreenshotAction = "take_screenshot"
)

type ComputerActionRequestGetCursorPositionAction string

const (
	ComputerActionRequestGetCursorPositionActionGetCursorPosition ComputerActionRequestGetCursorPositionAction = "get_cursor_position"
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

type CreateSessionRequestUseProxyGeolocationCity string

const (
	CreateSessionRequestUseProxyGeolocationCityACoruna                 CreateSessionRequestUseProxyGeolocationCity = "A_CORUNA"
	CreateSessionRequestUseProxyGeolocationCityAbidjan                 CreateSessionRequestUseProxyGeolocationCity = "ABIDJAN"
	CreateSessionRequestUseProxyGeolocationCityAbuDhabi                CreateSessionRequestUseProxyGeolocationCity = "ABU_DHABI"
	CreateSessionRequestUseProxyGeolocationCityAbuja                   CreateSessionRequestUseProxyGeolocationCity = "ABUJA"
	CreateSessionRequestUseProxyGeolocationCityAcapulcoDeJuarez        CreateSessionRequestUseProxyGeolocationCity = "ACAPULCO_DE JUAREZ"
	CreateSessionRequestUseProxyGeolocationCityAccra                   CreateSessionRequestUseProxyGeolocationCity = "ACCRA"
	CreateSessionRequestUseProxyGeolocationCityAdana                   CreateSessionRequestUseProxyGeolocationCity = "ADANA"
	CreateSessionRequestUseProxyGeolocationCityAdapazari               CreateSessionRequestUseProxyGeolocationCity = "ADAPAZARI"
	CreateSessionRequestUseProxyGeolocationCityAddisAbaba              CreateSessionRequestUseProxyGeolocationCity = "ADDIS_ABABA"
	CreateSessionRequestUseProxyGeolocationCityAdelaide                CreateSessionRequestUseProxyGeolocationCity = "ADELAIDE"
	CreateSessionRequestUseProxyGeolocationCityAfyonkarahisar          CreateSessionRequestUseProxyGeolocationCity = "AFYONKARAHISAR"
	CreateSessionRequestUseProxyGeolocationCityAgadir                  CreateSessionRequestUseProxyGeolocationCity = "AGADIR"
	CreateSessionRequestUseProxyGeolocationCityAguasLindasDeGoias      CreateSessionRequestUseProxyGeolocationCity = "AGUAS_LINDAS DE GOIAS"
	CreateSessionRequestUseProxyGeolocationCityAguascalientes          CreateSessionRequestUseProxyGeolocationCity = "AGUASCALIENTES"
	CreateSessionRequestUseProxyGeolocationCityAhmedabad               CreateSessionRequestUseProxyGeolocationCity = "AHMEDABAD"
	CreateSessionRequestUseProxyGeolocationCityAizawl                  CreateSessionRequestUseProxyGeolocationCity = "AIZAWL"
	CreateSessionRequestUseProxyGeolocationCityAjman                   CreateSessionRequestUseProxyGeolocationCity = "AJMAN"
	CreateSessionRequestUseProxyGeolocationCityAkron                   CreateSessionRequestUseProxyGeolocationCity = "AKRON"
	CreateSessionRequestUseProxyGeolocationCityAksaray                 CreateSessionRequestUseProxyGeolocationCity = "AKSARAY"
	CreateSessionRequestUseProxyGeolocationCityAlAinCity               CreateSessionRequestUseProxyGeolocationCity = "AL_AIN CITY"
	CreateSessionRequestUseProxyGeolocationCityAlMansurah              CreateSessionRequestUseProxyGeolocationCity = "AL_MANSURAH"
	CreateSessionRequestUseProxyGeolocationCityAlQatif                 CreateSessionRequestUseProxyGeolocationCity = "AL_QATIF"
	CreateSessionRequestUseProxyGeolocationCityAlajuela                CreateSessionRequestUseProxyGeolocationCity = "ALAJUELA"
	CreateSessionRequestUseProxyGeolocationCityAlbany                  CreateSessionRequestUseProxyGeolocationCity = "ALBANY"
	CreateSessionRequestUseProxyGeolocationCityAlbuquerque             CreateSessionRequestUseProxyGeolocationCity = "ALBUQUERQUE"
	CreateSessionRequestUseProxyGeolocationCityAlexandria              CreateSessionRequestUseProxyGeolocationCity = "ALEXANDRIA"
	CreateSessionRequestUseProxyGeolocationCityAlgiers                 CreateSessionRequestUseProxyGeolocationCity = "ALGIERS"
	CreateSessionRequestUseProxyGeolocationCityAlicante                CreateSessionRequestUseProxyGeolocationCity = "ALICANTE"
	CreateSessionRequestUseProxyGeolocationCityAlmada                  CreateSessionRequestUseProxyGeolocationCity = "ALMADA"
	CreateSessionRequestUseProxyGeolocationCityAlmaty                  CreateSessionRequestUseProxyGeolocationCity = "ALMATY"
	CreateSessionRequestUseProxyGeolocationCityAlmereStad              CreateSessionRequestUseProxyGeolocationCity = "ALMERE_STAD"
	CreateSessionRequestUseProxyGeolocationCityAlvorada                CreateSessionRequestUseProxyGeolocationCity = "ALVORADA"
	CreateSessionRequestUseProxyGeolocationCityAmadora                 CreateSessionRequestUseProxyGeolocationCity = "AMADORA"
	CreateSessionRequestUseProxyGeolocationCityAmasya                  CreateSessionRequestUseProxyGeolocationCity = "AMASYA"
	CreateSessionRequestUseProxyGeolocationCityAmbato                  CreateSessionRequestUseProxyGeolocationCity = "AMBATO"
	CreateSessionRequestUseProxyGeolocationCityAmericana               CreateSessionRequestUseProxyGeolocationCity = "AMERICANA"
	CreateSessionRequestUseProxyGeolocationCityAmman                   CreateSessionRequestUseProxyGeolocationCity = "AMMAN"
	CreateSessionRequestUseProxyGeolocationCityAmsterdam               CreateSessionRequestUseProxyGeolocationCity = "AMSTERDAM"
	CreateSessionRequestUseProxyGeolocationCityAnanindeua              CreateSessionRequestUseProxyGeolocationCity = "ANANINDEUA"
	CreateSessionRequestUseProxyGeolocationCityAnapolis                CreateSessionRequestUseProxyGeolocationCity = "ANAPOLIS"
	CreateSessionRequestUseProxyGeolocationCityAngelesCity             CreateSessionRequestUseProxyGeolocationCity = "ANGELES_CITY"
	CreateSessionRequestUseProxyGeolocationCityAngers                  CreateSessionRequestUseProxyGeolocationCity = "ANGERS"
	CreateSessionRequestUseProxyGeolocationCityAngraDosReis            CreateSessionRequestUseProxyGeolocationCity = "ANGRA_DOS REIS"
	CreateSessionRequestUseProxyGeolocationCityAnkara                  CreateSessionRequestUseProxyGeolocationCity = "ANKARA"
	CreateSessionRequestUseProxyGeolocationCityAntakya                 CreateSessionRequestUseProxyGeolocationCity = "ANTAKYA"
	CreateSessionRequestUseProxyGeolocationCityAntalya                 CreateSessionRequestUseProxyGeolocationCity = "ANTALYA"
	CreateSessionRequestUseProxyGeolocationCityAntananarivo            CreateSessionRequestUseProxyGeolocationCity = "ANTANANARIVO"
	CreateSessionRequestUseProxyGeolocationCityAntipoloCity            CreateSessionRequestUseProxyGeolocationCity = "ANTIPOLO_CITY"
	CreateSessionRequestUseProxyGeolocationCityAntofagasta             CreateSessionRequestUseProxyGeolocationCity = "ANTOFAGASTA"
	CreateSessionRequestUseProxyGeolocationCityAntwerp                 CreateSessionRequestUseProxyGeolocationCity = "ANTWERP"
	CreateSessionRequestUseProxyGeolocationCityAparecidaDeGoiania      CreateSessionRequestUseProxyGeolocationCity = "APARECIDA_DE GOIANIA"
	CreateSessionRequestUseProxyGeolocationCityApodaca                 CreateSessionRequestUseProxyGeolocationCity = "APODACA"
	CreateSessionRequestUseProxyGeolocationCityAracaju                 CreateSessionRequestUseProxyGeolocationCity = "ARACAJU"
	CreateSessionRequestUseProxyGeolocationCityAracatuba               CreateSessionRequestUseProxyGeolocationCity = "ARACATUBA"
	CreateSessionRequestUseProxyGeolocationCityArad                    CreateSessionRequestUseProxyGeolocationCity = "ARAD"
	CreateSessionRequestUseProxyGeolocationCityAraguaina               CreateSessionRequestUseProxyGeolocationCity = "ARAGUAINA"
	CreateSessionRequestUseProxyGeolocationCityArapiraca               CreateSessionRequestUseProxyGeolocationCity = "ARAPIRACA"
	CreateSessionRequestUseProxyGeolocationCityAraraquara              CreateSessionRequestUseProxyGeolocationCity = "ARARAQUARA"
	CreateSessionRequestUseProxyGeolocationCityArequipa                CreateSessionRequestUseProxyGeolocationCity = "AREQUIPA"
	CreateSessionRequestUseProxyGeolocationCityArica                   CreateSessionRequestUseProxyGeolocationCity = "ARICA"
	CreateSessionRequestUseProxyGeolocationCityArlington               CreateSessionRequestUseProxyGeolocationCity = "ARLINGTON"
	CreateSessionRequestUseProxyGeolocationCityAryanah                 CreateSessionRequestUseProxyGeolocationCity = "ARYANAH"
	CreateSessionRequestUseProxyGeolocationCityAstana                  CreateSessionRequestUseProxyGeolocationCity = "ASTANA"
	CreateSessionRequestUseProxyGeolocationCityAsuncion                CreateSessionRequestUseProxyGeolocationCity = "ASUNCION"
	CreateSessionRequestUseProxyGeolocationCityAsyut                   CreateSessionRequestUseProxyGeolocationCity = "ASYUT"
	CreateSessionRequestUseProxyGeolocationCityAtakum                  CreateSessionRequestUseProxyGeolocationCity = "ATAKUM"
	CreateSessionRequestUseProxyGeolocationCityAthens                  CreateSessionRequestUseProxyGeolocationCity = "ATHENS"
	CreateSessionRequestUseProxyGeolocationCityAtibaia                 CreateSessionRequestUseProxyGeolocationCity = "ATIBAIA"
	CreateSessionRequestUseProxyGeolocationCityAtlanta                 CreateSessionRequestUseProxyGeolocationCity = "ATLANTA"
	CreateSessionRequestUseProxyGeolocationCityAuburn                  CreateSessionRequestUseProxyGeolocationCity = "AUBURN"
	CreateSessionRequestUseProxyGeolocationCityAuckland                CreateSessionRequestUseProxyGeolocationCity = "AUCKLAND"
	CreateSessionRequestUseProxyGeolocationCityAurora                  CreateSessionRequestUseProxyGeolocationCity = "AURORA"
	CreateSessionRequestUseProxyGeolocationCityAustin                  CreateSessionRequestUseProxyGeolocationCity = "AUSTIN"
	CreateSessionRequestUseProxyGeolocationCityAvellaneda              CreateSessionRequestUseProxyGeolocationCity = "AVELLANEDA"
	CreateSessionRequestUseProxyGeolocationCityAydin                   CreateSessionRequestUseProxyGeolocationCity = "AYDIN"
	CreateSessionRequestUseProxyGeolocationCityAzcapotzalco            CreateSessionRequestUseProxyGeolocationCity = "AZCAPOTZALCO"
	CreateSessionRequestUseProxyGeolocationCityBacolodCity             CreateSessionRequestUseProxyGeolocationCity = "BACOLOD_CITY"
	CreateSessionRequestUseProxyGeolocationCityBacoor                  CreateSessionRequestUseProxyGeolocationCity = "BACOOR"
	CreateSessionRequestUseProxyGeolocationCityBaghdad                 CreateSessionRequestUseProxyGeolocationCity = "BAGHDAD"
	CreateSessionRequestUseProxyGeolocationCityBaguioCity              CreateSessionRequestUseProxyGeolocationCity = "BAGUIO_CITY"
	CreateSessionRequestUseProxyGeolocationCityBahiaBlanca             CreateSessionRequestUseProxyGeolocationCity = "BAHIA_BLANCA"
	CreateSessionRequestUseProxyGeolocationCityBakersfield             CreateSessionRequestUseProxyGeolocationCity = "BAKERSFIELD"
	CreateSessionRequestUseProxyGeolocationCityBaku                    CreateSessionRequestUseProxyGeolocationCity = "BAKU"
	CreateSessionRequestUseProxyGeolocationCityBalikesir               CreateSessionRequestUseProxyGeolocationCity = "BALIKESIR"
	CreateSessionRequestUseProxyGeolocationCityBalikpapan              CreateSessionRequestUseProxyGeolocationCity = "BALIKPAPAN"
	CreateSessionRequestUseProxyGeolocationCityBalnearioCamboriu       CreateSessionRequestUseProxyGeolocationCity = "BALNEARIO_CAMBORIU"
	CreateSessionRequestUseProxyGeolocationCityBaltimore               CreateSessionRequestUseProxyGeolocationCity = "BALTIMORE"
	CreateSessionRequestUseProxyGeolocationCityBandarLampung           CreateSessionRequestUseProxyGeolocationCity = "BANDAR_LAMPUNG"
	CreateSessionRequestUseProxyGeolocationCityBandarSeriBegawan       CreateSessionRequestUseProxyGeolocationCity = "BANDAR_SERI BEGAWAN"
	CreateSessionRequestUseProxyGeolocationCityBandung                 CreateSessionRequestUseProxyGeolocationCity = "BANDUNG"
	CreateSessionRequestUseProxyGeolocationCityBangkok                 CreateSessionRequestUseProxyGeolocationCity = "BANGKOK"
	CreateSessionRequestUseProxyGeolocationCityBanjaLuka               CreateSessionRequestUseProxyGeolocationCity = "BANJA_LUKA"
	CreateSessionRequestUseProxyGeolocationCityBanjarmasin             CreateSessionRequestUseProxyGeolocationCity = "BANJARMASIN"
	CreateSessionRequestUseProxyGeolocationCityBarcelona               CreateSessionRequestUseProxyGeolocationCity = "BARCELONA"
	CreateSessionRequestUseProxyGeolocationCityBari                    CreateSessionRequestUseProxyGeolocationCity = "BARI"
	CreateSessionRequestUseProxyGeolocationCityBarquisimeto            CreateSessionRequestUseProxyGeolocationCity = "BARQUISIMETO"
	CreateSessionRequestUseProxyGeolocationCityBarraMansa              CreateSessionRequestUseProxyGeolocationCity = "BARRA_MANSA"
	CreateSessionRequestUseProxyGeolocationCityBarranquilla            CreateSessionRequestUseProxyGeolocationCity = "BARRANQUILLA"
	CreateSessionRequestUseProxyGeolocationCityBarueri                 CreateSessionRequestUseProxyGeolocationCity = "BARUERI"
	CreateSessionRequestUseProxyGeolocationCityBatam                   CreateSessionRequestUseProxyGeolocationCity = "BATAM"
	CreateSessionRequestUseProxyGeolocationCityBatangas                CreateSessionRequestUseProxyGeolocationCity = "BATANGAS"
	CreateSessionRequestUseProxyGeolocationCityBatman                  CreateSessionRequestUseProxyGeolocationCity = "BATMAN"
	CreateSessionRequestUseProxyGeolocationCityBatnaCity               CreateSessionRequestUseProxyGeolocationCity = "BATNA_CITY"
	CreateSessionRequestUseProxyGeolocationCityBatonRouge              CreateSessionRequestUseProxyGeolocationCity = "BATON_ROUGE"
	CreateSessionRequestUseProxyGeolocationCityBatumi                  CreateSessionRequestUseProxyGeolocationCity = "BATUMI"
	CreateSessionRequestUseProxyGeolocationCityBauru                   CreateSessionRequestUseProxyGeolocationCity = "BAURU"
	CreateSessionRequestUseProxyGeolocationCityBeirut                  CreateSessionRequestUseProxyGeolocationCity = "BEIRUT"
	CreateSessionRequestUseProxyGeolocationCityBejaia                  CreateSessionRequestUseProxyGeolocationCity = "BEJAIA"
	CreateSessionRequestUseProxyGeolocationCityBekasi                  CreateSessionRequestUseProxyGeolocationCity = "BEKASI"
	CreateSessionRequestUseProxyGeolocationCityBelem                   CreateSessionRequestUseProxyGeolocationCity = "BELEM"
	CreateSessionRequestUseProxyGeolocationCityBelfast                 CreateSessionRequestUseProxyGeolocationCity = "BELFAST"
	CreateSessionRequestUseProxyGeolocationCityBelfordRoxo             CreateSessionRequestUseProxyGeolocationCity = "BELFORD_ROXO"
	CreateSessionRequestUseProxyGeolocationCityBelgrade                CreateSessionRequestUseProxyGeolocationCity = "BELGRADE"
	CreateSessionRequestUseProxyGeolocationCityBeloHorizonte           CreateSessionRequestUseProxyGeolocationCity = "BELO_HORIZONTE"
	CreateSessionRequestUseProxyGeolocationCityBengaluru               CreateSessionRequestUseProxyGeolocationCity = "BENGALURU"
	CreateSessionRequestUseProxyGeolocationCityBeniMellal              CreateSessionRequestUseProxyGeolocationCity = "BENI_MELLAL"
	CreateSessionRequestUseProxyGeolocationCityBerazategui             CreateSessionRequestUseProxyGeolocationCity = "BERAZATEGUI"
	CreateSessionRequestUseProxyGeolocationCityBern                    CreateSessionRequestUseProxyGeolocationCity = "BERN"
	CreateSessionRequestUseProxyGeolocationCityBetim                   CreateSessionRequestUseProxyGeolocationCity = "BETIM"
	CreateSessionRequestUseProxyGeolocationCityBharatpur               CreateSessionRequestUseProxyGeolocationCity = "BHARATPUR"
	CreateSessionRequestUseProxyGeolocationCityBhopal                  CreateSessionRequestUseProxyGeolocationCity = "BHOPAL"
	CreateSessionRequestUseProxyGeolocationCityBhubaneswar             CreateSessionRequestUseProxyGeolocationCity = "BHUBANESWAR"
	CreateSessionRequestUseProxyGeolocationCityBialystok               CreateSessionRequestUseProxyGeolocationCity = "BIALYSTOK"
	CreateSessionRequestUseProxyGeolocationCityBienHoa                 CreateSessionRequestUseProxyGeolocationCity = "BIEN_HOA"
	CreateSessionRequestUseProxyGeolocationCityBilbao                  CreateSessionRequestUseProxyGeolocationCity = "BILBAO"
	CreateSessionRequestUseProxyGeolocationCityBilecik                 CreateSessionRequestUseProxyGeolocationCity = "BILECIK"
	CreateSessionRequestUseProxyGeolocationCityBiratnagar              CreateSessionRequestUseProxyGeolocationCity = "BIRATNAGAR"
	CreateSessionRequestUseProxyGeolocationCityBirmingham              CreateSessionRequestUseProxyGeolocationCity = "BIRMINGHAM"
	CreateSessionRequestUseProxyGeolocationCityBishkek                 CreateSessionRequestUseProxyGeolocationCity = "BISHKEK"
	CreateSessionRequestUseProxyGeolocationCityBizerte                 CreateSessionRequestUseProxyGeolocationCity = "BIZERTE"
	CreateSessionRequestUseProxyGeolocationCityBlida                   CreateSessionRequestUseProxyGeolocationCity = "BLIDA"
	CreateSessionRequestUseProxyGeolocationCityBloemfontein            CreateSessionRequestUseProxyGeolocationCity = "BLOEMFONTEIN"
	CreateSessionRequestUseProxyGeolocationCityBloomington             CreateSessionRequestUseProxyGeolocationCity = "BLOOMINGTON"
	CreateSessionRequestUseProxyGeolocationCityBlumenau                CreateSessionRequestUseProxyGeolocationCity = "BLUMENAU"
	CreateSessionRequestUseProxyGeolocationCityBoaVista                CreateSessionRequestUseProxyGeolocationCity = "BOA_VISTA"
	CreateSessionRequestUseProxyGeolocationCityBochum                  CreateSessionRequestUseProxyGeolocationCity = "BOCHUM"
	CreateSessionRequestUseProxyGeolocationCityBogor                   CreateSessionRequestUseProxyGeolocationCity = "BOGOR"
	CreateSessionRequestUseProxyGeolocationCityBogota                  CreateSessionRequestUseProxyGeolocationCity = "BOGOTA"
	CreateSessionRequestUseProxyGeolocationCityBoise                   CreateSessionRequestUseProxyGeolocationCity = "BOISE"
	CreateSessionRequestUseProxyGeolocationCityBoksburg                CreateSessionRequestUseProxyGeolocationCity = "BOKSBURG"
	CreateSessionRequestUseProxyGeolocationCityBologna                 CreateSessionRequestUseProxyGeolocationCity = "BOLOGNA"
	CreateSessionRequestUseProxyGeolocationCityBolu                    CreateSessionRequestUseProxyGeolocationCity = "BOLU"
	CreateSessionRequestUseProxyGeolocationCityBordeaux                CreateSessionRequestUseProxyGeolocationCity = "BORDEAUX"
	CreateSessionRequestUseProxyGeolocationCityBoston                  CreateSessionRequestUseProxyGeolocationCity = "BOSTON"
	CreateSessionRequestUseProxyGeolocationCityBotucatu                CreateSessionRequestUseProxyGeolocationCity = "BOTUCATU"
	CreateSessionRequestUseProxyGeolocationCityBradford                CreateSessionRequestUseProxyGeolocationCity = "BRADFORD"
	CreateSessionRequestUseProxyGeolocationCityBraga                   CreateSessionRequestUseProxyGeolocationCity = "BRAGA"
	CreateSessionRequestUseProxyGeolocationCityBragancaPaulista        CreateSessionRequestUseProxyGeolocationCity = "BRAGANCA_PAULISTA"
	CreateSessionRequestUseProxyGeolocationCityBrampton                CreateSessionRequestUseProxyGeolocationCity = "BRAMPTON"
	CreateSessionRequestUseProxyGeolocationCityBrasilia                CreateSessionRequestUseProxyGeolocationCity = "BRASILIA"
	CreateSessionRequestUseProxyGeolocationCityBrasov                  CreateSessionRequestUseProxyGeolocationCity = "BRASOV"
	CreateSessionRequestUseProxyGeolocationCityBratislava              CreateSessionRequestUseProxyGeolocationCity = "BRATISLAVA"
	CreateSessionRequestUseProxyGeolocationCityBremen                  CreateSessionRequestUseProxyGeolocationCity = "BREMEN"
	CreateSessionRequestUseProxyGeolocationCityBrescia                 CreateSessionRequestUseProxyGeolocationCity = "BRESCIA"
	CreateSessionRequestUseProxyGeolocationCityBrest                   CreateSessionRequestUseProxyGeolocationCity = "BREST"
	CreateSessionRequestUseProxyGeolocationCityBridgetown              CreateSessionRequestUseProxyGeolocationCity = "BRIDGETOWN"
	CreateSessionRequestUseProxyGeolocationCityBrisbane                CreateSessionRequestUseProxyGeolocationCity = "BRISBANE"
	CreateSessionRequestUseProxyGeolocationCityBristol                 CreateSessionRequestUseProxyGeolocationCity = "BRISTOL"
	CreateSessionRequestUseProxyGeolocationCityBrno                    CreateSessionRequestUseProxyGeolocationCity = "BRNO"
	CreateSessionRequestUseProxyGeolocationCityBrooklyn                CreateSessionRequestUseProxyGeolocationCity = "BROOKLYN"
	CreateSessionRequestUseProxyGeolocationCityBrussels                CreateSessionRequestUseProxyGeolocationCity = "BRUSSELS"
	CreateSessionRequestUseProxyGeolocationCityBucaramanga             CreateSessionRequestUseProxyGeolocationCity = "BUCARAMANGA"
	CreateSessionRequestUseProxyGeolocationCityBucharest               CreateSessionRequestUseProxyGeolocationCity = "BUCHAREST"
	CreateSessionRequestUseProxyGeolocationCityBudapest                CreateSessionRequestUseProxyGeolocationCity = "BUDAPEST"
	CreateSessionRequestUseProxyGeolocationCityBuenosAires             CreateSessionRequestUseProxyGeolocationCity = "BUENOS_AIRES"
	CreateSessionRequestUseProxyGeolocationCityBuffalo                 CreateSessionRequestUseProxyGeolocationCity = "BUFFALO"
	CreateSessionRequestUseProxyGeolocationCityBukGu                   CreateSessionRequestUseProxyGeolocationCity = "BUK_GU"
	CreateSessionRequestUseProxyGeolocationCityBukhara                 CreateSessionRequestUseProxyGeolocationCity = "BUKHARA"
	CreateSessionRequestUseProxyGeolocationCityBurgas                  CreateSessionRequestUseProxyGeolocationCity = "BURGAS"
	CreateSessionRequestUseProxyGeolocationCityBurnaby                 CreateSessionRequestUseProxyGeolocationCity = "BURNABY"
	CreateSessionRequestUseProxyGeolocationCityBursa                   CreateSessionRequestUseProxyGeolocationCity = "BURSA"
	CreateSessionRequestUseProxyGeolocationCityButuan                  CreateSessionRequestUseProxyGeolocationCity = "BUTUAN"
	CreateSessionRequestUseProxyGeolocationCityBydgoszcz               CreateSessionRequestUseProxyGeolocationCity = "BYDGOSZCZ"
	CreateSessionRequestUseProxyGeolocationCityCabanatuanCity          CreateSessionRequestUseProxyGeolocationCity = "CABANATUAN_CITY"
	CreateSessionRequestUseProxyGeolocationCityCaboFrio                CreateSessionRequestUseProxyGeolocationCity = "CABO_FRIO"
	CreateSessionRequestUseProxyGeolocationCityCabuyao                 CreateSessionRequestUseProxyGeolocationCity = "CABUYAO"
	CreateSessionRequestUseProxyGeolocationCityCachoeiroDeItapemirim   CreateSessionRequestUseProxyGeolocationCity = "CACHOEIRO_DE ITAPEMIRIM"
	CreateSessionRequestUseProxyGeolocationCityCagayanDeOro            CreateSessionRequestUseProxyGeolocationCity = "CAGAYAN_DE ORO"
	CreateSessionRequestUseProxyGeolocationCityCagliari                CreateSessionRequestUseProxyGeolocationCity = "CAGLIARI"
	CreateSessionRequestUseProxyGeolocationCityCairo                   CreateSessionRequestUseProxyGeolocationCity = "CAIRO"
	CreateSessionRequestUseProxyGeolocationCityCalamba                 CreateSessionRequestUseProxyGeolocationCity = "CALAMBA"
	CreateSessionRequestUseProxyGeolocationCityCalgary                 CreateSessionRequestUseProxyGeolocationCity = "CALGARY"
	CreateSessionRequestUseProxyGeolocationCityCaloocanCity            CreateSessionRequestUseProxyGeolocationCity = "CALOOCAN_CITY"
	CreateSessionRequestUseProxyGeolocationCityCamacari                CreateSessionRequestUseProxyGeolocationCity = "CAMACARI"
	CreateSessionRequestUseProxyGeolocationCityCamaragibe              CreateSessionRequestUseProxyGeolocationCity = "CAMARAGIBE"
	CreateSessionRequestUseProxyGeolocationCityCampeche                CreateSessionRequestUseProxyGeolocationCity = "CAMPECHE"
	CreateSessionRequestUseProxyGeolocationCityCampinaGrande           CreateSessionRequestUseProxyGeolocationCity = "CAMPINA_GRANDE"
	CreateSessionRequestUseProxyGeolocationCityCampinas                CreateSessionRequestUseProxyGeolocationCity = "CAMPINAS"
	CreateSessionRequestUseProxyGeolocationCityCampoGrande             CreateSessionRequestUseProxyGeolocationCity = "CAMPO_GRANDE"
	CreateSessionRequestUseProxyGeolocationCityCampoLargo              CreateSessionRequestUseProxyGeolocationCity = "CAMPO_LARGO"
	CreateSessionRequestUseProxyGeolocationCityCamposDosGoytacazes     CreateSessionRequestUseProxyGeolocationCity = "CAMPOS_DOS GOYTACAZES"
	CreateSessionRequestUseProxyGeolocationCityCanTho                  CreateSessionRequestUseProxyGeolocationCity = "CAN_THO"
	CreateSessionRequestUseProxyGeolocationCityCanoas                  CreateSessionRequestUseProxyGeolocationCity = "CANOAS"
	CreateSessionRequestUseProxyGeolocationCityCanton                  CreateSessionRequestUseProxyGeolocationCity = "CANTON"
	CreateSessionRequestUseProxyGeolocationCityCapeTown                CreateSessionRequestUseProxyGeolocationCity = "CAPE_TOWN"
	CreateSessionRequestUseProxyGeolocationCityCaracas                 CreateSessionRequestUseProxyGeolocationCity = "CARACAS"
	CreateSessionRequestUseProxyGeolocationCityCaraguatatuba           CreateSessionRequestUseProxyGeolocationCity = "CARAGUATATUBA"
	CreateSessionRequestUseProxyGeolocationCityCarapicuiba             CreateSessionRequestUseProxyGeolocationCity = "CARAPICUIBA"
	CreateSessionRequestUseProxyGeolocationCityCardiff                 CreateSessionRequestUseProxyGeolocationCity = "CARDIFF"
	CreateSessionRequestUseProxyGeolocationCityCariacica               CreateSessionRequestUseProxyGeolocationCity = "CARIACICA"
	CreateSessionRequestUseProxyGeolocationCityCarmona                 CreateSessionRequestUseProxyGeolocationCity = "CARMONA"
	CreateSessionRequestUseProxyGeolocationCityCartagena               CreateSessionRequestUseProxyGeolocationCity = "CARTAGENA"
	CreateSessionRequestUseProxyGeolocationCityCaruaru                 CreateSessionRequestUseProxyGeolocationCity = "CARUARU"
	CreateSessionRequestUseProxyGeolocationCityCasablanca              CreateSessionRequestUseProxyGeolocationCity = "CASABLANCA"
	CreateSessionRequestUseProxyGeolocationCityCascavel                CreateSessionRequestUseProxyGeolocationCity = "CASCAVEL"
	CreateSessionRequestUseProxyGeolocationCityCaseros                 CreateSessionRequestUseProxyGeolocationCity = "CASEROS"
	CreateSessionRequestUseProxyGeolocationCityCastanhal               CreateSessionRequestUseProxyGeolocationCity = "CASTANHAL"
	CreateSessionRequestUseProxyGeolocationCityCastries                CreateSessionRequestUseProxyGeolocationCity = "CASTRIES"
	CreateSessionRequestUseProxyGeolocationCityCatalao                 CreateSessionRequestUseProxyGeolocationCity = "CATALAO"
	CreateSessionRequestUseProxyGeolocationCityCatamarca               CreateSessionRequestUseProxyGeolocationCity = "CATAMARCA"
	CreateSessionRequestUseProxyGeolocationCityCatanduva               CreateSessionRequestUseProxyGeolocationCity = "CATANDUVA"
	CreateSessionRequestUseProxyGeolocationCityCatania                 CreateSessionRequestUseProxyGeolocationCity = "CATANIA"
	CreateSessionRequestUseProxyGeolocationCityCaucaia                 CreateSessionRequestUseProxyGeolocationCity = "CAUCAIA"
	CreateSessionRequestUseProxyGeolocationCityCaxiasDoSul             CreateSessionRequestUseProxyGeolocationCity = "CAXIAS_DO SUL"
	CreateSessionRequestUseProxyGeolocationCityCebuCity                CreateSessionRequestUseProxyGeolocationCity = "CEBU_CITY"
	CreateSessionRequestUseProxyGeolocationCityCentral                 CreateSessionRequestUseProxyGeolocationCity = "CENTRAL"
	CreateSessionRequestUseProxyGeolocationCityCentro                  CreateSessionRequestUseProxyGeolocationCity = "CENTRO"
	CreateSessionRequestUseProxyGeolocationCityCenturion               CreateSessionRequestUseProxyGeolocationCity = "CENTURION"
	CreateSessionRequestUseProxyGeolocationCityChaguanas               CreateSessionRequestUseProxyGeolocationCity = "CHAGUANAS"
	CreateSessionRequestUseProxyGeolocationCityChandigarh              CreateSessionRequestUseProxyGeolocationCity = "CHANDIGARH"
	CreateSessionRequestUseProxyGeolocationCityChandler                CreateSessionRequestUseProxyGeolocationCity = "CHANDLER"
	CreateSessionRequestUseProxyGeolocationCityChangHua                CreateSessionRequestUseProxyGeolocationCity = "CHANG_HUA"
	CreateSessionRequestUseProxyGeolocationCityChapeco                 CreateSessionRequestUseProxyGeolocationCity = "CHAPECO"
	CreateSessionRequestUseProxyGeolocationCityCharleston              CreateSessionRequestUseProxyGeolocationCity = "CHARLESTON"
	CreateSessionRequestUseProxyGeolocationCityCharlotte               CreateSessionRequestUseProxyGeolocationCity = "CHARLOTTE"
	CreateSessionRequestUseProxyGeolocationCityChelyabinsk             CreateSessionRequestUseProxyGeolocationCity = "CHELYABINSK"
	CreateSessionRequestUseProxyGeolocationCityChennai                 CreateSessionRequestUseProxyGeolocationCity = "CHENNAI"
	CreateSessionRequestUseProxyGeolocationCityCherkasy                CreateSessionRequestUseProxyGeolocationCity = "CHERKASY"
	CreateSessionRequestUseProxyGeolocationCityChernivtsi              CreateSessionRequestUseProxyGeolocationCity = "CHERNIVTSI"
	CreateSessionRequestUseProxyGeolocationCityChia                    CreateSessionRequestUseProxyGeolocationCity = "CHIA"
	CreateSessionRequestUseProxyGeolocationCityChiangMai               CreateSessionRequestUseProxyGeolocationCity = "CHIANG_MAI"
	CreateSessionRequestUseProxyGeolocationCityChiclayo                CreateSessionRequestUseProxyGeolocationCity = "CHICLAYO"
	CreateSessionRequestUseProxyGeolocationCityChihuahuaCity           CreateSessionRequestUseProxyGeolocationCity = "CHIHUAHUA_CITY"
	CreateSessionRequestUseProxyGeolocationCityChimbote                CreateSessionRequestUseProxyGeolocationCity = "CHIMBOTE"
	CreateSessionRequestUseProxyGeolocationCityChisinau                CreateSessionRequestUseProxyGeolocationCity = "CHISINAU"
	CreateSessionRequestUseProxyGeolocationCityChittagong              CreateSessionRequestUseProxyGeolocationCity = "CHITTAGONG"
	CreateSessionRequestUseProxyGeolocationCityChristchurch            CreateSessionRequestUseProxyGeolocationCity = "CHRISTCHURCH"
	CreateSessionRequestUseProxyGeolocationCityCincinnati              CreateSessionRequestUseProxyGeolocationCity = "CINCINNATI"
	CreateSessionRequestUseProxyGeolocationCityCirebon                 CreateSessionRequestUseProxyGeolocationCity = "CIREBON"
	CreateSessionRequestUseProxyGeolocationCityCityOfMuntinlupa        CreateSessionRequestUseProxyGeolocationCity = "CITY_OF MUNTINLUPA"
	CreateSessionRequestUseProxyGeolocationCityCiudadDelEste           CreateSessionRequestUseProxyGeolocationCity = "CIUDAD_DEL ESTE"
	CreateSessionRequestUseProxyGeolocationCityCiudadGuayana           CreateSessionRequestUseProxyGeolocationCity = "CIUDAD_GUAYANA"
	CreateSessionRequestUseProxyGeolocationCityCiudadJuarez            CreateSessionRequestUseProxyGeolocationCity = "CIUDAD_JUAREZ"
	CreateSessionRequestUseProxyGeolocationCityCiudadNezahualcoyotl    CreateSessionRequestUseProxyGeolocationCity = "CIUDAD_NEZAHUALCOYOTL"
	CreateSessionRequestUseProxyGeolocationCityCiudadObregon           CreateSessionRequestUseProxyGeolocationCity = "CIUDAD_OBREGON"
	CreateSessionRequestUseProxyGeolocationCityCleveland               CreateSessionRequestUseProxyGeolocationCity = "CLEVELAND"
	CreateSessionRequestUseProxyGeolocationCityClujNapoca              CreateSessionRequestUseProxyGeolocationCity = "CLUJ_NAPOCA"
	CreateSessionRequestUseProxyGeolocationCityCochabamba              CreateSessionRequestUseProxyGeolocationCity = "COCHABAMBA"
	CreateSessionRequestUseProxyGeolocationCityCoimbatore              CreateSessionRequestUseProxyGeolocationCity = "COIMBATORE"
	CreateSessionRequestUseProxyGeolocationCityCoimbra                 CreateSessionRequestUseProxyGeolocationCity = "COIMBRA"
	CreateSessionRequestUseProxyGeolocationCityCologne                 CreateSessionRequestUseProxyGeolocationCity = "COLOGNE"
	CreateSessionRequestUseProxyGeolocationCityColombo                 CreateSessionRequestUseProxyGeolocationCity = "COLOMBO"
	CreateSessionRequestUseProxyGeolocationCityColoradoSprings         CreateSessionRequestUseProxyGeolocationCity = "COLORADO_SPRINGS"
	CreateSessionRequestUseProxyGeolocationCityColumbia                CreateSessionRequestUseProxyGeolocationCity = "COLUMBIA"
	CreateSessionRequestUseProxyGeolocationCityColumbus                CreateSessionRequestUseProxyGeolocationCity = "COLUMBUS"
	CreateSessionRequestUseProxyGeolocationCityComodoroRivadavia       CreateSessionRequestUseProxyGeolocationCity = "COMODORO_RIVADAVIA"
	CreateSessionRequestUseProxyGeolocationCityConcepcion              CreateSessionRequestUseProxyGeolocationCity = "CONCEPCION"
	CreateSessionRequestUseProxyGeolocationCityConcord                 CreateSessionRequestUseProxyGeolocationCity = "CONCORD"
	CreateSessionRequestUseProxyGeolocationCityConstanta               CreateSessionRequestUseProxyGeolocationCity = "CONSTANTA"
	CreateSessionRequestUseProxyGeolocationCityConstantine             CreateSessionRequestUseProxyGeolocationCity = "CONSTANTINE"
	CreateSessionRequestUseProxyGeolocationCityContagem                CreateSessionRequestUseProxyGeolocationCity = "CONTAGEM"
	CreateSessionRequestUseProxyGeolocationCityCopenhagen              CreateSessionRequestUseProxyGeolocationCity = "COPENHAGEN"
	CreateSessionRequestUseProxyGeolocationCityCordoba                 CreateSessionRequestUseProxyGeolocationCity = "CORDOBA"
	CreateSessionRequestUseProxyGeolocationCityCorrientes              CreateSessionRequestUseProxyGeolocationCity = "CORRIENTES"
	CreateSessionRequestUseProxyGeolocationCityCorum                   CreateSessionRequestUseProxyGeolocationCity = "CORUM"
	CreateSessionRequestUseProxyGeolocationCityCotia                   CreateSessionRequestUseProxyGeolocationCity = "COTIA"
	CreateSessionRequestUseProxyGeolocationCityCoventry                CreateSessionRequestUseProxyGeolocationCity = "COVENTRY"
	CreateSessionRequestUseProxyGeolocationCityCraiova                 CreateSessionRequestUseProxyGeolocationCity = "CRAIOVA"
	CreateSessionRequestUseProxyGeolocationCityCriciuma                CreateSessionRequestUseProxyGeolocationCity = "CRICIUMA"
	CreateSessionRequestUseProxyGeolocationCityCroydon                 CreateSessionRequestUseProxyGeolocationCity = "CROYDON"
	CreateSessionRequestUseProxyGeolocationCityCuautitlanIzcalli       CreateSessionRequestUseProxyGeolocationCity = "CUAUTITLAN_IZCALLI"
	CreateSessionRequestUseProxyGeolocationCityCucuta                  CreateSessionRequestUseProxyGeolocationCity = "CUCUTA"
	CreateSessionRequestUseProxyGeolocationCityCuenca                  CreateSessionRequestUseProxyGeolocationCity = "CUENCA"
	CreateSessionRequestUseProxyGeolocationCityCuernavaca              CreateSessionRequestUseProxyGeolocationCity = "CUERNAVACA"
	CreateSessionRequestUseProxyGeolocationCityCuiaba                  CreateSessionRequestUseProxyGeolocationCity = "CUIABA"
	CreateSessionRequestUseProxyGeolocationCityCuliacan                CreateSessionRequestUseProxyGeolocationCity = "CULIACAN"
	CreateSessionRequestUseProxyGeolocationCityCuritiba                CreateSessionRequestUseProxyGeolocationCity = "CURITIBA"
	CreateSessionRequestUseProxyGeolocationCityCusco                   CreateSessionRequestUseProxyGeolocationCity = "CUSCO"
	CreateSessionRequestUseProxyGeolocationCityDaNang                  CreateSessionRequestUseProxyGeolocationCity = "DA_NANG"
	CreateSessionRequestUseProxyGeolocationCityDagupan                 CreateSessionRequestUseProxyGeolocationCity = "DAGUPAN"
	CreateSessionRequestUseProxyGeolocationCityDakar                   CreateSessionRequestUseProxyGeolocationCity = "DAKAR"
	CreateSessionRequestUseProxyGeolocationCityDallas                  CreateSessionRequestUseProxyGeolocationCity = "DALLAS"
	CreateSessionRequestUseProxyGeolocationCityDamietta                CreateSessionRequestUseProxyGeolocationCity = "DAMIETTA"
	CreateSessionRequestUseProxyGeolocationCityDammam                  CreateSessionRequestUseProxyGeolocationCity = "DAMMAM"
	CreateSessionRequestUseProxyGeolocationCityDarEsSalaam             CreateSessionRequestUseProxyGeolocationCity = "DAR_ES SALAAM"
	CreateSessionRequestUseProxyGeolocationCityDasmarinas              CreateSessionRequestUseProxyGeolocationCity = "DASMARINAS"
	CreateSessionRequestUseProxyGeolocationCityDavaoCity               CreateSessionRequestUseProxyGeolocationCity = "DAVAO_CITY"
	CreateSessionRequestUseProxyGeolocationCityDayton                  CreateSessionRequestUseProxyGeolocationCity = "DAYTON"
	CreateSessionRequestUseProxyGeolocationCityDebrecen                CreateSessionRequestUseProxyGeolocationCity = "DEBRECEN"
	CreateSessionRequestUseProxyGeolocationCityDecatur                 CreateSessionRequestUseProxyGeolocationCity = "DECATUR"
	CreateSessionRequestUseProxyGeolocationCityDehradun                CreateSessionRequestUseProxyGeolocationCity = "DEHRADUN"
	CreateSessionRequestUseProxyGeolocationCityDelhi                   CreateSessionRequestUseProxyGeolocationCity = "DELHI"
	CreateSessionRequestUseProxyGeolocationCityDenizli                 CreateSessionRequestUseProxyGeolocationCity = "DENIZLI"
	CreateSessionRequestUseProxyGeolocationCityDenpasar                CreateSessionRequestUseProxyGeolocationCity = "DENPASAR"
	CreateSessionRequestUseProxyGeolocationCityDenver                  CreateSessionRequestUseProxyGeolocationCity = "DENVER"
	CreateSessionRequestUseProxyGeolocationCityDepok                   CreateSessionRequestUseProxyGeolocationCity = "DEPOK"
	CreateSessionRequestUseProxyGeolocationCityDerby                   CreateSessionRequestUseProxyGeolocationCity = "DERBY"
	CreateSessionRequestUseProxyGeolocationCityDetroit                 CreateSessionRequestUseProxyGeolocationCity = "DETROIT"
	CreateSessionRequestUseProxyGeolocationCityDhaka                   CreateSessionRequestUseProxyGeolocationCity = "DHAKA"
	CreateSessionRequestUseProxyGeolocationCityDiadema                 CreateSessionRequestUseProxyGeolocationCity = "DIADEMA"
	CreateSessionRequestUseProxyGeolocationCityDivinopolis             CreateSessionRequestUseProxyGeolocationCity = "DIVINOPOLIS"
	CreateSessionRequestUseProxyGeolocationCityDiyarbakir              CreateSessionRequestUseProxyGeolocationCity = "DIYARBAKIR"
	CreateSessionRequestUseProxyGeolocationCityDjelfa                  CreateSessionRequestUseProxyGeolocationCity = "DJELFA"
	CreateSessionRequestUseProxyGeolocationCityDnipro                  CreateSessionRequestUseProxyGeolocationCity = "DNIPRO"
	CreateSessionRequestUseProxyGeolocationCityDoha                    CreateSessionRequestUseProxyGeolocationCity = "DOHA"
	CreateSessionRequestUseProxyGeolocationCityDortmund                CreateSessionRequestUseProxyGeolocationCity = "DORTMUND"
	CreateSessionRequestUseProxyGeolocationCityDourados                CreateSessionRequestUseProxyGeolocationCity = "DOURADOS"
	CreateSessionRequestUseProxyGeolocationCityDresden                 CreateSessionRequestUseProxyGeolocationCity = "DRESDEN"
	CreateSessionRequestUseProxyGeolocationCityDubai                   CreateSessionRequestUseProxyGeolocationCity = "DUBAI"
	CreateSessionRequestUseProxyGeolocationCityDublin                  CreateSessionRequestUseProxyGeolocationCity = "DUBLIN"
	CreateSessionRequestUseProxyGeolocationCityDuezce                  CreateSessionRequestUseProxyGeolocationCity = "DUEZCE"
	CreateSessionRequestUseProxyGeolocationCityDuisburg                CreateSessionRequestUseProxyGeolocationCity = "DUISBURG"
	CreateSessionRequestUseProxyGeolocationCityDuqueDeCaxias           CreateSessionRequestUseProxyGeolocationCity = "DUQUE_DE CAXIAS"
	CreateSessionRequestUseProxyGeolocationCityDurango                 CreateSessionRequestUseProxyGeolocationCity = "DURANGO"
	CreateSessionRequestUseProxyGeolocationCityDurban                  CreateSessionRequestUseProxyGeolocationCity = "DURBAN"
	CreateSessionRequestUseProxyGeolocationCityDusseldorf              CreateSessionRequestUseProxyGeolocationCity = "DUSSELDORF"
	CreateSessionRequestUseProxyGeolocationCityEcatepec                CreateSessionRequestUseProxyGeolocationCity = "ECATEPEC"
	CreateSessionRequestUseProxyGeolocationCityEdinburgh               CreateSessionRequestUseProxyGeolocationCity = "EDINBURGH"
	CreateSessionRequestUseProxyGeolocationCityEdirne                  CreateSessionRequestUseProxyGeolocationCity = "EDIRNE"
	CreateSessionRequestUseProxyGeolocationCityEdmonton                CreateSessionRequestUseProxyGeolocationCity = "EDMONTON"
	CreateSessionRequestUseProxyGeolocationCityElJadida                CreateSessionRequestUseProxyGeolocationCity = "EL_JADIDA"
	CreateSessionRequestUseProxyGeolocationCityElPaso                  CreateSessionRequestUseProxyGeolocationCity = "EL_PASO"
	CreateSessionRequestUseProxyGeolocationCityElazig                  CreateSessionRequestUseProxyGeolocationCity = "ELAZIG"
	CreateSessionRequestUseProxyGeolocationCityEmbu                    CreateSessionRequestUseProxyGeolocationCity = "EMBU"
	CreateSessionRequestUseProxyGeolocationCityEnsenada                CreateSessionRequestUseProxyGeolocationCity = "ENSENADA"
	CreateSessionRequestUseProxyGeolocationCityErbil                   CreateSessionRequestUseProxyGeolocationCity = "ERBIL"
	CreateSessionRequestUseProxyGeolocationCityErzurum                 CreateSessionRequestUseProxyGeolocationCity = "ERZURUM"
	CreateSessionRequestUseProxyGeolocationCityEskisehir               CreateSessionRequestUseProxyGeolocationCity = "ESKISEHIR"
	CreateSessionRequestUseProxyGeolocationCityEspoo                   CreateSessionRequestUseProxyGeolocationCity = "ESPOO"
	CreateSessionRequestUseProxyGeolocationCityEssen                   CreateSessionRequestUseProxyGeolocationCity = "ESSEN"
	CreateSessionRequestUseProxyGeolocationCityFaisalabad              CreateSessionRequestUseProxyGeolocationCity = "FAISALABAD"
	CreateSessionRequestUseProxyGeolocationCityFayetteville            CreateSessionRequestUseProxyGeolocationCity = "FAYETTEVILLE"
	CreateSessionRequestUseProxyGeolocationCityFazendaRioGrande        CreateSessionRequestUseProxyGeolocationCity = "FAZENDA_RIO GRANDE"
	CreateSessionRequestUseProxyGeolocationCityFeiraDeSantana          CreateSessionRequestUseProxyGeolocationCity = "FEIRA_DE SANTANA"
	CreateSessionRequestUseProxyGeolocationCityFes                     CreateSessionRequestUseProxyGeolocationCity = "FES"
	CreateSessionRequestUseProxyGeolocationCityFlorence                CreateSessionRequestUseProxyGeolocationCity = "FLORENCE"
	CreateSessionRequestUseProxyGeolocationCityFlorencioVarela         CreateSessionRequestUseProxyGeolocationCity = "FLORENCIO_VARELA"
	CreateSessionRequestUseProxyGeolocationCityFlorianopolis           CreateSessionRequestUseProxyGeolocationCity = "FLORIANOPOLIS"
	CreateSessionRequestUseProxyGeolocationCityFontana                 CreateSessionRequestUseProxyGeolocationCity = "FONTANA"
	CreateSessionRequestUseProxyGeolocationCityFormosa                 CreateSessionRequestUseProxyGeolocationCity = "FORMOSA"
	CreateSessionRequestUseProxyGeolocationCityFortLauderdale          CreateSessionRequestUseProxyGeolocationCity = "FORT_LAUDERDALE"
	CreateSessionRequestUseProxyGeolocationCityFortWayne               CreateSessionRequestUseProxyGeolocationCity = "FORT_WAYNE"
	CreateSessionRequestUseProxyGeolocationCityFortWorth               CreateSessionRequestUseProxyGeolocationCity = "FORT_WORTH"
	CreateSessionRequestUseProxyGeolocationCityFortaleza               CreateSessionRequestUseProxyGeolocationCity = "FORTALEZA"
	CreateSessionRequestUseProxyGeolocationCityFozDoIguacu             CreateSessionRequestUseProxyGeolocationCity = "FOZ_DO IGUACU"
	CreateSessionRequestUseProxyGeolocationCityFranca                  CreateSessionRequestUseProxyGeolocationCity = "FRANCA"
	CreateSessionRequestUseProxyGeolocationCityFranciscoMorato         CreateSessionRequestUseProxyGeolocationCity = "FRANCISCO_MORATO"
	CreateSessionRequestUseProxyGeolocationCityFrancoDaRocha           CreateSessionRequestUseProxyGeolocationCity = "FRANCO_DA ROCHA"
	CreateSessionRequestUseProxyGeolocationCityFrankfurtAmMain         CreateSessionRequestUseProxyGeolocationCity = "FRANKFURT_AM MAIN"
	CreateSessionRequestUseProxyGeolocationCityFredericksburg          CreateSessionRequestUseProxyGeolocationCity = "FREDERICKSBURG"
	CreateSessionRequestUseProxyGeolocationCityFresno                  CreateSessionRequestUseProxyGeolocationCity = "FRESNO"
	CreateSessionRequestUseProxyGeolocationCityFunchal                 CreateSessionRequestUseProxyGeolocationCity = "FUNCHAL"
	CreateSessionRequestUseProxyGeolocationCityGaborone                CreateSessionRequestUseProxyGeolocationCity = "GABORONE"
	CreateSessionRequestUseProxyGeolocationCityGainesville             CreateSessionRequestUseProxyGeolocationCity = "GAINESVILLE"
	CreateSessionRequestUseProxyGeolocationCityGalati                  CreateSessionRequestUseProxyGeolocationCity = "GALATI"
	CreateSessionRequestUseProxyGeolocationCityGangnamGu               CreateSessionRequestUseProxyGeolocationCity = "GANGNAM_GU"
	CreateSessionRequestUseProxyGeolocationCityGaranhuns               CreateSessionRequestUseProxyGeolocationCity = "GARANHUNS"
	CreateSessionRequestUseProxyGeolocationCityGatineau                CreateSessionRequestUseProxyGeolocationCity = "GATINEAU"
	CreateSessionRequestUseProxyGeolocationCityGaziantep               CreateSessionRequestUseProxyGeolocationCity = "GAZIANTEP"
	CreateSessionRequestUseProxyGeolocationCityGdansk                  CreateSessionRequestUseProxyGeolocationCity = "GDANSK"
	CreateSessionRequestUseProxyGeolocationCityGdynia                  CreateSessionRequestUseProxyGeolocationCity = "GDYNIA"
	CreateSessionRequestUseProxyGeolocationCityGeneralTrias            CreateSessionRequestUseProxyGeolocationCity = "GENERAL_TRIAS"
	CreateSessionRequestUseProxyGeolocationCityGeneva                  CreateSessionRequestUseProxyGeolocationCity = "GENEVA"
	CreateSessionRequestUseProxyGeolocationCityGenoa                   CreateSessionRequestUseProxyGeolocationCity = "GENOA"
	CreateSessionRequestUseProxyGeolocationCityGeorgeTown              CreateSessionRequestUseProxyGeolocationCity = "GEORGE_TOWN"
	CreateSessionRequestUseProxyGeolocationCityGeorgetown              CreateSessionRequestUseProxyGeolocationCity = "GEORGETOWN"
	CreateSessionRequestUseProxyGeolocationCityGhaziabad               CreateSessionRequestUseProxyGeolocationCity = "GHAZIABAD"
	CreateSessionRequestUseProxyGeolocationCityGhent                   CreateSessionRequestUseProxyGeolocationCity = "GHENT"
	CreateSessionRequestUseProxyGeolocationCityGijon                   CreateSessionRequestUseProxyGeolocationCity = "GIJON"
	CreateSessionRequestUseProxyGeolocationCityGiresun                 CreateSessionRequestUseProxyGeolocationCity = "GIRESUN"
	CreateSessionRequestUseProxyGeolocationCityGiza                    CreateSessionRequestUseProxyGeolocationCity = "GIZA"
	CreateSessionRequestUseProxyGeolocationCityGlasgow                 CreateSessionRequestUseProxyGeolocationCity = "GLASGOW"
	CreateSessionRequestUseProxyGeolocationCityGlendale                CreateSessionRequestUseProxyGeolocationCity = "GLENDALE"
	CreateSessionRequestUseProxyGeolocationCityGliwice                 CreateSessionRequestUseProxyGeolocationCity = "GLIWICE"
	CreateSessionRequestUseProxyGeolocationCityGoiania                 CreateSessionRequestUseProxyGeolocationCity = "GOIANIA"
	CreateSessionRequestUseProxyGeolocationCityGomel                   CreateSessionRequestUseProxyGeolocationCity = "GOMEL"
	CreateSessionRequestUseProxyGeolocationCityGothenburg              CreateSessionRequestUseProxyGeolocationCity = "GOTHENBURG"
	CreateSessionRequestUseProxyGeolocationCityGovernadorValadares     CreateSessionRequestUseProxyGeolocationCity = "GOVERNADOR_VALADARES"
	CreateSessionRequestUseProxyGeolocationCityGoyangSi                CreateSessionRequestUseProxyGeolocationCity = "GOYANG_SI"
	CreateSessionRequestUseProxyGeolocationCityGranada                 CreateSessionRequestUseProxyGeolocationCity = "GRANADA"
	CreateSessionRequestUseProxyGeolocationCityGrandRapids             CreateSessionRequestUseProxyGeolocationCity = "GRAND_RAPIDS"
	CreateSessionRequestUseProxyGeolocationCityGravatai                CreateSessionRequestUseProxyGeolocationCity = "GRAVATAI"
	CreateSessionRequestUseProxyGeolocationCityGraz                    CreateSessionRequestUseProxyGeolocationCity = "GRAZ"
	CreateSessionRequestUseProxyGeolocationCityGreensboro              CreateSessionRequestUseProxyGeolocationCity = "GREENSBORO"
	CreateSessionRequestUseProxyGeolocationCityGreenville              CreateSessionRequestUseProxyGeolocationCity = "GREENVILLE"
	CreateSessionRequestUseProxyGeolocationCityGuadalajara             CreateSessionRequestUseProxyGeolocationCity = "GUADALAJARA"
	CreateSessionRequestUseProxyGeolocationCityGuadalupe               CreateSessionRequestUseProxyGeolocationCity = "GUADALUPE"
	CreateSessionRequestUseProxyGeolocationCityGuangzhou               CreateSessionRequestUseProxyGeolocationCity = "GUANGZHOU"
	CreateSessionRequestUseProxyGeolocationCityGuarapuava              CreateSessionRequestUseProxyGeolocationCity = "GUARAPUAVA"
	CreateSessionRequestUseProxyGeolocationCityGuaratingueta           CreateSessionRequestUseProxyGeolocationCity = "GUARATINGUETA"
	CreateSessionRequestUseProxyGeolocationCityGuaruja                 CreateSessionRequestUseProxyGeolocationCity = "GUARUJA"
	CreateSessionRequestUseProxyGeolocationCityGuarulhos               CreateSessionRequestUseProxyGeolocationCity = "GUARULHOS"
	CreateSessionRequestUseProxyGeolocationCityGuatemalaCity           CreateSessionRequestUseProxyGeolocationCity = "GUATEMALA_CITY"
	CreateSessionRequestUseProxyGeolocationCityGuayaquil               CreateSessionRequestUseProxyGeolocationCity = "GUAYAQUIL"
	CreateSessionRequestUseProxyGeolocationCityGujranwala              CreateSessionRequestUseProxyGeolocationCity = "GUJRANWALA"
	CreateSessionRequestUseProxyGeolocationCityGurugram                CreateSessionRequestUseProxyGeolocationCity = "GURUGRAM"
	CreateSessionRequestUseProxyGeolocationCityGustavoAdolfoMadero     CreateSessionRequestUseProxyGeolocationCity = "GUSTAVO_ADOLFO MADERO"
	CreateSessionRequestUseProxyGeolocationCityGuwahati                CreateSessionRequestUseProxyGeolocationCity = "GUWAHATI"
	CreateSessionRequestUseProxyGeolocationCityGwanakGu                CreateSessionRequestUseProxyGeolocationCity = "GWANAK_GU"
	CreateSessionRequestUseProxyGeolocationCityHackney                 CreateSessionRequestUseProxyGeolocationCity = "HACKNEY"
	CreateSessionRequestUseProxyGeolocationCityHaifa                   CreateSessionRequestUseProxyGeolocationCity = "HAIFA"
	CreateSessionRequestUseProxyGeolocationCityHaiphong                CreateSessionRequestUseProxyGeolocationCity = "HAIPHONG"
	CreateSessionRequestUseProxyGeolocationCityHamburg                 CreateSessionRequestUseProxyGeolocationCity = "HAMBURG"
	CreateSessionRequestUseProxyGeolocationCityHamilton                CreateSessionRequestUseProxyGeolocationCity = "HAMILTON"
	CreateSessionRequestUseProxyGeolocationCityHanoi                   CreateSessionRequestUseProxyGeolocationCity = "HANOI"
	CreateSessionRequestUseProxyGeolocationCityHanover                 CreateSessionRequestUseProxyGeolocationCity = "HANOVER"
	CreateSessionRequestUseProxyGeolocationCityHarare                  CreateSessionRequestUseProxyGeolocationCity = "HARARE"
	CreateSessionRequestUseProxyGeolocationCityHavana                  CreateSessionRequestUseProxyGeolocationCity = "HAVANA"
	CreateSessionRequestUseProxyGeolocationCityHelsinki                CreateSessionRequestUseProxyGeolocationCity = "HELSINKI"
	CreateSessionRequestUseProxyGeolocationCityHenderson               CreateSessionRequestUseProxyGeolocationCity = "HENDERSON"
	CreateSessionRequestUseProxyGeolocationCityHeredia                 CreateSessionRequestUseProxyGeolocationCity = "HEREDIA"
	CreateSessionRequestUseProxyGeolocationCityHermosillo              CreateSessionRequestUseProxyGeolocationCity = "HERMOSILLO"
	CreateSessionRequestUseProxyGeolocationCityHialeah                 CreateSessionRequestUseProxyGeolocationCity = "HIALEAH"
	CreateSessionRequestUseProxyGeolocationCityHoChiMinhCity           CreateSessionRequestUseProxyGeolocationCity = "HO_CHI MINH CITY"
	CreateSessionRequestUseProxyGeolocationCityHollywood               CreateSessionRequestUseProxyGeolocationCity = "HOLLYWOOD"
	CreateSessionRequestUseProxyGeolocationCityHolon                   CreateSessionRequestUseProxyGeolocationCity = "HOLON"
	CreateSessionRequestUseProxyGeolocationCityHonolulu                CreateSessionRequestUseProxyGeolocationCity = "HONOLULU"
	CreateSessionRequestUseProxyGeolocationCityHortolandia             CreateSessionRequestUseProxyGeolocationCity = "HORTOLANDIA"
	CreateSessionRequestUseProxyGeolocationCityHrodna                  CreateSessionRequestUseProxyGeolocationCity = "HRODNA"
	CreateSessionRequestUseProxyGeolocationCityHsinchu                 CreateSessionRequestUseProxyGeolocationCity = "HSINCHU"
	CreateSessionRequestUseProxyGeolocationCityHuancayo                CreateSessionRequestUseProxyGeolocationCity = "HUANCAYO"
	CreateSessionRequestUseProxyGeolocationCityHuanuco                 CreateSessionRequestUseProxyGeolocationCity = "HUANUCO"
	CreateSessionRequestUseProxyGeolocationCityHull                    CreateSessionRequestUseProxyGeolocationCity = "HULL"
	CreateSessionRequestUseProxyGeolocationCityHurlingham              CreateSessionRequestUseProxyGeolocationCity = "HURLINGHAM"
	CreateSessionRequestUseProxyGeolocationCityHyderabad               CreateSessionRequestUseProxyGeolocationCity = "HYDERABAD"
	CreateSessionRequestUseProxyGeolocationCityIasi                    CreateSessionRequestUseProxyGeolocationCity = "IASI"
	CreateSessionRequestUseProxyGeolocationCityIbague                  CreateSessionRequestUseProxyGeolocationCity = "IBAGUE"
	CreateSessionRequestUseProxyGeolocationCityIca                     CreateSessionRequestUseProxyGeolocationCity = "ICA"
	CreateSessionRequestUseProxyGeolocationCityIlam                    CreateSessionRequestUseProxyGeolocationCity = "ILAM"
	CreateSessionRequestUseProxyGeolocationCityIlford                  CreateSessionRequestUseProxyGeolocationCity = "ILFORD"
	CreateSessionRequestUseProxyGeolocationCityIligan                  CreateSessionRequestUseProxyGeolocationCity = "ILIGAN"
	CreateSessionRequestUseProxyGeolocationCityIloiloCity              CreateSessionRequestUseProxyGeolocationCity = "ILOILO_CITY"
	CreateSessionRequestUseProxyGeolocationCityImperatriz              CreateSessionRequestUseProxyGeolocationCity = "IMPERATRIZ"
	CreateSessionRequestUseProxyGeolocationCityImus                    CreateSessionRequestUseProxyGeolocationCity = "IMUS"
	CreateSessionRequestUseProxyGeolocationCityIncheon                 CreateSessionRequestUseProxyGeolocationCity = "INCHEON"
	CreateSessionRequestUseProxyGeolocationCityIndaiatuba              CreateSessionRequestUseProxyGeolocationCity = "INDAIATUBA"
	CreateSessionRequestUseProxyGeolocationCityIndianapolis            CreateSessionRequestUseProxyGeolocationCity = "INDIANAPOLIS"
	CreateSessionRequestUseProxyGeolocationCityIndore                  CreateSessionRequestUseProxyGeolocationCity = "INDORE"
	CreateSessionRequestUseProxyGeolocationCityIpatinga                CreateSessionRequestUseProxyGeolocationCity = "IPATINGA"
	CreateSessionRequestUseProxyGeolocationCityIpoh                    CreateSessionRequestUseProxyGeolocationCity = "IPOH"
	CreateSessionRequestUseProxyGeolocationCityIquique                 CreateSessionRequestUseProxyGeolocationCity = "IQUIQUE"
	CreateSessionRequestUseProxyGeolocationCityIrvine                  CreateSessionRequestUseProxyGeolocationCity = "IRVINE"
	CreateSessionRequestUseProxyGeolocationCityIsidroCasanova          CreateSessionRequestUseProxyGeolocationCity = "ISIDRO_CASANOVA"
	CreateSessionRequestUseProxyGeolocationCityIslamabad               CreateSessionRequestUseProxyGeolocationCity = "ISLAMABAD"
	CreateSessionRequestUseProxyGeolocationCityIslington               CreateSessionRequestUseProxyGeolocationCity = "ISLINGTON"
	CreateSessionRequestUseProxyGeolocationCityIsmailia                CreateSessionRequestUseProxyGeolocationCity = "ISMAILIA"
	CreateSessionRequestUseProxyGeolocationCityIsparta                 CreateSessionRequestUseProxyGeolocationCity = "ISPARTA"
	CreateSessionRequestUseProxyGeolocationCityIstanbul                CreateSessionRequestUseProxyGeolocationCity = "ISTANBUL"
	CreateSessionRequestUseProxyGeolocationCityItaborai                CreateSessionRequestUseProxyGeolocationCity = "ITABORAI"
	CreateSessionRequestUseProxyGeolocationCityItabuna                 CreateSessionRequestUseProxyGeolocationCity = "ITABUNA"
	CreateSessionRequestUseProxyGeolocationCityItajai                  CreateSessionRequestUseProxyGeolocationCity = "ITAJAI"
	CreateSessionRequestUseProxyGeolocationCityItanhaem                CreateSessionRequestUseProxyGeolocationCity = "ITANHAEM"
	CreateSessionRequestUseProxyGeolocationCityItapevi                 CreateSessionRequestUseProxyGeolocationCity = "ITAPEVI"
	CreateSessionRequestUseProxyGeolocationCityItaquaquecetuba         CreateSessionRequestUseProxyGeolocationCity = "ITAQUAQUECETUBA"
	CreateSessionRequestUseProxyGeolocationCityItuzaingo               CreateSessionRequestUseProxyGeolocationCity = "ITUZAINGO"
	CreateSessionRequestUseProxyGeolocationCityIzmir                   CreateSessionRequestUseProxyGeolocationCity = "IZMIR"
	CreateSessionRequestUseProxyGeolocationCityIztapalapa              CreateSessionRequestUseProxyGeolocationCity = "IZTAPALAPA"
	CreateSessionRequestUseProxyGeolocationCityJaboataoDosGuararapes   CreateSessionRequestUseProxyGeolocationCity = "JABOATAO_DOS GUARARAPES"
	CreateSessionRequestUseProxyGeolocationCityJacarei                 CreateSessionRequestUseProxyGeolocationCity = "JACAREI"
	CreateSessionRequestUseProxyGeolocationCityJackson                 CreateSessionRequestUseProxyGeolocationCity = "JACKSON"
	CreateSessionRequestUseProxyGeolocationCityJacksonville            CreateSessionRequestUseProxyGeolocationCity = "JACKSONVILLE"
	CreateSessionRequestUseProxyGeolocationCityJaipur                  CreateSessionRequestUseProxyGeolocationCity = "JAIPUR"
	CreateSessionRequestUseProxyGeolocationCityJakarta                 CreateSessionRequestUseProxyGeolocationCity = "JAKARTA"
	CreateSessionRequestUseProxyGeolocationCityJaraguaDoSul            CreateSessionRequestUseProxyGeolocationCity = "JARAGUA_DO SUL"
	CreateSessionRequestUseProxyGeolocationCityJau                     CreateSessionRequestUseProxyGeolocationCity = "JAU"
	CreateSessionRequestUseProxyGeolocationCityJeddah                  CreateSessionRequestUseProxyGeolocationCity = "JEDDAH"
	CreateSessionRequestUseProxyGeolocationCityJember                  CreateSessionRequestUseProxyGeolocationCity = "JEMBER"
	CreateSessionRequestUseProxyGeolocationCityJerusalem               CreateSessionRequestUseProxyGeolocationCity = "JERUSALEM"
	CreateSessionRequestUseProxyGeolocationCityJoaoMonlevade           CreateSessionRequestUseProxyGeolocationCity = "JOAO_MONLEVADE"
	CreateSessionRequestUseProxyGeolocationCityJoaoPessoa              CreateSessionRequestUseProxyGeolocationCity = "JOAO_PESSOA"
	CreateSessionRequestUseProxyGeolocationCityJodhpur                 CreateSessionRequestUseProxyGeolocationCity = "JODHPUR"
	CreateSessionRequestUseProxyGeolocationCityJohannesburg            CreateSessionRequestUseProxyGeolocationCity = "JOHANNESBURG"
	CreateSessionRequestUseProxyGeolocationCityJohorBahru              CreateSessionRequestUseProxyGeolocationCity = "JOHOR_BAHRU"
	CreateSessionRequestUseProxyGeolocationCityJoinville               CreateSessionRequestUseProxyGeolocationCity = "JOINVILLE"
	CreateSessionRequestUseProxyGeolocationCityJoseCPaz                CreateSessionRequestUseProxyGeolocationCity = "JOSE_C PAZ"
	CreateSessionRequestUseProxyGeolocationCityJoseMariaEzeiza         CreateSessionRequestUseProxyGeolocationCity = "JOSE_MARIA EZEIZA"
	CreateSessionRequestUseProxyGeolocationCityJuarez                  CreateSessionRequestUseProxyGeolocationCity = "JUAREZ"
	CreateSessionRequestUseProxyGeolocationCityJuazeiroDoNorte         CreateSessionRequestUseProxyGeolocationCity = "JUAZEIRO_DO NORTE"
	CreateSessionRequestUseProxyGeolocationCityJuizDeFora              CreateSessionRequestUseProxyGeolocationCity = "JUIZ_DE FORA"
	CreateSessionRequestUseProxyGeolocationCityJundiai                 CreateSessionRequestUseProxyGeolocationCity = "JUNDIAI"
	CreateSessionRequestUseProxyGeolocationCityKahramanmaras           CreateSessionRequestUseProxyGeolocationCity = "KAHRAMANMARAS"
	CreateSessionRequestUseProxyGeolocationCityKampala                 CreateSessionRequestUseProxyGeolocationCity = "KAMPALA"
	CreateSessionRequestUseProxyGeolocationCityKanpur                  CreateSessionRequestUseProxyGeolocationCity = "KANPUR"
	CreateSessionRequestUseProxyGeolocationCityKansasCity              CreateSessionRequestUseProxyGeolocationCity = "KANSAS_CITY"
	CreateSessionRequestUseProxyGeolocationCityKaohsiungCity           CreateSessionRequestUseProxyGeolocationCity = "KAOHSIUNG_CITY"
	CreateSessionRequestUseProxyGeolocationCityKarabuk                 CreateSessionRequestUseProxyGeolocationCity = "KARABUK"
	CreateSessionRequestUseProxyGeolocationCityKarachi                 CreateSessionRequestUseProxyGeolocationCity = "KARACHI"
	CreateSessionRequestUseProxyGeolocationCityKarlsruhe               CreateSessionRequestUseProxyGeolocationCity = "KARLSRUHE"
	CreateSessionRequestUseProxyGeolocationCityKarnal                  CreateSessionRequestUseProxyGeolocationCity = "KARNAL"
	CreateSessionRequestUseProxyGeolocationCityKaski                   CreateSessionRequestUseProxyGeolocationCity = "KASKI"
	CreateSessionRequestUseProxyGeolocationCityKastamonu               CreateSessionRequestUseProxyGeolocationCity = "KASTAMONU"
	CreateSessionRequestUseProxyGeolocationCityKathmandu               CreateSessionRequestUseProxyGeolocationCity = "KATHMANDU"
	CreateSessionRequestUseProxyGeolocationCityKatowice                CreateSessionRequestUseProxyGeolocationCity = "KATOWICE"
	CreateSessionRequestUseProxyGeolocationCityKatsina                 CreateSessionRequestUseProxyGeolocationCity = "KATSINA"
	CreateSessionRequestUseProxyGeolocationCityKaty                    CreateSessionRequestUseProxyGeolocationCity = "KATY"
	CreateSessionRequestUseProxyGeolocationCityKaunas                  CreateSessionRequestUseProxyGeolocationCity = "KAUNAS"
	CreateSessionRequestUseProxyGeolocationCityKayseri                 CreateSessionRequestUseProxyGeolocationCity = "KAYSERI"
	CreateSessionRequestUseProxyGeolocationCityKazan                   CreateSessionRequestUseProxyGeolocationCity = "KAZAN"
	CreateSessionRequestUseProxyGeolocationCityKecskemet               CreateSessionRequestUseProxyGeolocationCity = "KECSKEMET"
	CreateSessionRequestUseProxyGeolocationCityKediri                  CreateSessionRequestUseProxyGeolocationCity = "KEDIRI"
	CreateSessionRequestUseProxyGeolocationCityKenitra                 CreateSessionRequestUseProxyGeolocationCity = "KENITRA"
	CreateSessionRequestUseProxyGeolocationCityKharkiv                 CreateSessionRequestUseProxyGeolocationCity = "KHARKIV"
	CreateSessionRequestUseProxyGeolocationCityKhmelnytskyi            CreateSessionRequestUseProxyGeolocationCity = "KHMELNYTSKYI"
	CreateSessionRequestUseProxyGeolocationCityKhonKaen                CreateSessionRequestUseProxyGeolocationCity = "KHON_KAEN"
	CreateSessionRequestUseProxyGeolocationCityKielce                  CreateSessionRequestUseProxyGeolocationCity = "KIELCE"
	CreateSessionRequestUseProxyGeolocationCityKigali                  CreateSessionRequestUseProxyGeolocationCity = "KIGALI"
	CreateSessionRequestUseProxyGeolocationCityKingston                CreateSessionRequestUseProxyGeolocationCity = "KINGSTON"
	CreateSessionRequestUseProxyGeolocationCityKirklareli              CreateSessionRequestUseProxyGeolocationCity = "KIRKLARELI"
	CreateSessionRequestUseProxyGeolocationCityKissimmee               CreateSessionRequestUseProxyGeolocationCity = "KISSIMMEE"
	CreateSessionRequestUseProxyGeolocationCityKitchener               CreateSessionRequestUseProxyGeolocationCity = "KITCHENER"
	CreateSessionRequestUseProxyGeolocationCityKlaipeda                CreateSessionRequestUseProxyGeolocationCity = "KLAIPEDA"
	CreateSessionRequestUseProxyGeolocationCityKnoxville               CreateSessionRequestUseProxyGeolocationCity = "KNOXVILLE"
	CreateSessionRequestUseProxyGeolocationCityKochi                   CreateSessionRequestUseProxyGeolocationCity = "KOCHI"
	CreateSessionRequestUseProxyGeolocationCityKolkata                 CreateSessionRequestUseProxyGeolocationCity = "KOLKATA"
	CreateSessionRequestUseProxyGeolocationCityKollam                  CreateSessionRequestUseProxyGeolocationCity = "KOLLAM"
	CreateSessionRequestUseProxyGeolocationCityKonya                   CreateSessionRequestUseProxyGeolocationCity = "KONYA"
	CreateSessionRequestUseProxyGeolocationCityKosekoy                 CreateSessionRequestUseProxyGeolocationCity = "KOSEKOY"
	CreateSessionRequestUseProxyGeolocationCityKosice                  CreateSessionRequestUseProxyGeolocationCity = "KOSICE"
	CreateSessionRequestUseProxyGeolocationCityKotaKinabalu            CreateSessionRequestUseProxyGeolocationCity = "KOTA_KINABALU"
	CreateSessionRequestUseProxyGeolocationCityKozhikode               CreateSessionRequestUseProxyGeolocationCity = "KOZHIKODE"
	CreateSessionRequestUseProxyGeolocationCityKrakow                  CreateSessionRequestUseProxyGeolocationCity = "KRAKOW"
	CreateSessionRequestUseProxyGeolocationCityKrasnodar               CreateSessionRequestUseProxyGeolocationCity = "KRASNODAR"
	CreateSessionRequestUseProxyGeolocationCityKryvyiRih               CreateSessionRequestUseProxyGeolocationCity = "KRYVYI_RIH"
	CreateSessionRequestUseProxyGeolocationCityKualaLumpur             CreateSessionRequestUseProxyGeolocationCity = "KUALA_LUMPUR"
	CreateSessionRequestUseProxyGeolocationCityKuching                 CreateSessionRequestUseProxyGeolocationCity = "KUCHING"
	CreateSessionRequestUseProxyGeolocationCityKutahya                 CreateSessionRequestUseProxyGeolocationCity = "KUTAHYA"
	CreateSessionRequestUseProxyGeolocationCityKutaisi                 CreateSessionRequestUseProxyGeolocationCity = "KUTAISI"
	CreateSessionRequestUseProxyGeolocationCityKuwaitCity              CreateSessionRequestUseProxyGeolocationCity = "KUWAIT_CITY"
	CreateSessionRequestUseProxyGeolocationCityKyiv                    CreateSessionRequestUseProxyGeolocationCity = "KYIV"
	CreateSessionRequestUseProxyGeolocationCityLaPaz                   CreateSessionRequestUseProxyGeolocationCity = "LA_PAZ"
	CreateSessionRequestUseProxyGeolocationCityLaPlata                 CreateSessionRequestUseProxyGeolocationCity = "LA_PLATA"
	CreateSessionRequestUseProxyGeolocationCityLaRioja                 CreateSessionRequestUseProxyGeolocationCity = "LA_RIOJA"
	CreateSessionRequestUseProxyGeolocationCityLaSerena                CreateSessionRequestUseProxyGeolocationCity = "LA_SERENA"
	CreateSessionRequestUseProxyGeolocationCityLafayette               CreateSessionRequestUseProxyGeolocationCity = "LAFAYETTE"
	CreateSessionRequestUseProxyGeolocationCityLaferrere               CreateSessionRequestUseProxyGeolocationCity = "LAFERRERE"
	CreateSessionRequestUseProxyGeolocationCityLages                   CreateSessionRequestUseProxyGeolocationCity = "LAGES"
	CreateSessionRequestUseProxyGeolocationCityLagos                   CreateSessionRequestUseProxyGeolocationCity = "LAGOS"
	CreateSessionRequestUseProxyGeolocationCityLahore                  CreateSessionRequestUseProxyGeolocationCity = "LAHORE"
	CreateSessionRequestUseProxyGeolocationCityLahug                   CreateSessionRequestUseProxyGeolocationCity = "LAHUG"
	CreateSessionRequestUseProxyGeolocationCityLakeWorth               CreateSessionRequestUseProxyGeolocationCity = "LAKE_WORTH"
	CreateSessionRequestUseProxyGeolocationCityLakeland                CreateSessionRequestUseProxyGeolocationCity = "LAKELAND"
	CreateSessionRequestUseProxyGeolocationCityLancaster               CreateSessionRequestUseProxyGeolocationCity = "LANCASTER"
	CreateSessionRequestUseProxyGeolocationCityLanus                   CreateSessionRequestUseProxyGeolocationCity = "LANUS"
	CreateSessionRequestUseProxyGeolocationCityLasPalmasDeGranCanaria  CreateSessionRequestUseProxyGeolocationCity = "LAS_PALMAS DE GRAN CANARIA"
	CreateSessionRequestUseProxyGeolocationCityLasPinas                CreateSessionRequestUseProxyGeolocationCity = "LAS_PINAS"
	CreateSessionRequestUseProxyGeolocationCityLasVegas                CreateSessionRequestUseProxyGeolocationCity = "LAS_VEGAS"
	CreateSessionRequestUseProxyGeolocationCityLausanne                CreateSessionRequestUseProxyGeolocationCity = "LAUSANNE"
	CreateSessionRequestUseProxyGeolocationCityLaval                   CreateSessionRequestUseProxyGeolocationCity = "LAVAL"
	CreateSessionRequestUseProxyGeolocationCityLawrenceville           CreateSessionRequestUseProxyGeolocationCity = "LAWRENCEVILLE"
	CreateSessionRequestUseProxyGeolocationCityLeMans                  CreateSessionRequestUseProxyGeolocationCity = "LE_MANS"
	CreateSessionRequestUseProxyGeolocationCityLeeds                   CreateSessionRequestUseProxyGeolocationCity = "LEEDS"
	CreateSessionRequestUseProxyGeolocationCityLeicester               CreateSessionRequestUseProxyGeolocationCity = "LEICESTER"
	CreateSessionRequestUseProxyGeolocationCityLeipzig                 CreateSessionRequestUseProxyGeolocationCity = "LEIPZIG"
	CreateSessionRequestUseProxyGeolocationCityLeon                    CreateSessionRequestUseProxyGeolocationCity = "LEON"
	CreateSessionRequestUseProxyGeolocationCityLexington               CreateSessionRequestUseProxyGeolocationCity = "LEXINGTON"
	CreateSessionRequestUseProxyGeolocationCityLibreville              CreateSessionRequestUseProxyGeolocationCity = "LIBREVILLE"
	CreateSessionRequestUseProxyGeolocationCityLiege                   CreateSessionRequestUseProxyGeolocationCity = "LIEGE"
	CreateSessionRequestUseProxyGeolocationCityLille                   CreateSessionRequestUseProxyGeolocationCity = "LILLE"
	CreateSessionRequestUseProxyGeolocationCityLima                    CreateSessionRequestUseProxyGeolocationCity = "LIMA"
	CreateSessionRequestUseProxyGeolocationCityLimassol                CreateSessionRequestUseProxyGeolocationCity = "LIMASSOL"
	CreateSessionRequestUseProxyGeolocationCityLimeira                 CreateSessionRequestUseProxyGeolocationCity = "LIMEIRA"
	CreateSessionRequestUseProxyGeolocationCityLincoln                 CreateSessionRequestUseProxyGeolocationCity = "LINCOLN"
	CreateSessionRequestUseProxyGeolocationCityLinhares                CreateSessionRequestUseProxyGeolocationCity = "LINHARES"
	CreateSessionRequestUseProxyGeolocationCityLipaCity                CreateSessionRequestUseProxyGeolocationCity = "LIPA_CITY"
	CreateSessionRequestUseProxyGeolocationCityLisbon                  CreateSessionRequestUseProxyGeolocationCity = "LISBON"
	CreateSessionRequestUseProxyGeolocationCityLiverpool               CreateSessionRequestUseProxyGeolocationCity = "LIVERPOOL"
	CreateSessionRequestUseProxyGeolocationCityLjubljana               CreateSessionRequestUseProxyGeolocationCity = "LJUBLJANA"
	CreateSessionRequestUseProxyGeolocationCityLodz                    CreateSessionRequestUseProxyGeolocationCity = "LODZ"
	CreateSessionRequestUseProxyGeolocationCityLoja                    CreateSessionRequestUseProxyGeolocationCity = "LOJA"
	CreateSessionRequestUseProxyGeolocationCityLomasDeZamora           CreateSessionRequestUseProxyGeolocationCity = "LOMAS_DE ZAMORA"
	CreateSessionRequestUseProxyGeolocationCityLome                    CreateSessionRequestUseProxyGeolocationCity = "LOME"
	CreateSessionRequestUseProxyGeolocationCityLondrina                CreateSessionRequestUseProxyGeolocationCity = "LONDRINA"
	CreateSessionRequestUseProxyGeolocationCityLongBeach               CreateSessionRequestUseProxyGeolocationCity = "LONG_BEACH"
	CreateSessionRequestUseProxyGeolocationCityLongueuil               CreateSessionRequestUseProxyGeolocationCity = "LONGUEUIL"
	CreateSessionRequestUseProxyGeolocationCityLouisville              CreateSessionRequestUseProxyGeolocationCity = "LOUISVILLE"
	CreateSessionRequestUseProxyGeolocationCityLuanda                  CreateSessionRequestUseProxyGeolocationCity = "LUANDA"
	CreateSessionRequestUseProxyGeolocationCityLublin                  CreateSessionRequestUseProxyGeolocationCity = "LUBLIN"
	CreateSessionRequestUseProxyGeolocationCityLucenaCity              CreateSessionRequestUseProxyGeolocationCity = "LUCENA_CITY"
	CreateSessionRequestUseProxyGeolocationCityLucknow                 CreateSessionRequestUseProxyGeolocationCity = "LUCKNOW"
	CreateSessionRequestUseProxyGeolocationCityLudhiana                CreateSessionRequestUseProxyGeolocationCity = "LUDHIANA"
	CreateSessionRequestUseProxyGeolocationCityLusaka                  CreateSessionRequestUseProxyGeolocationCity = "LUSAKA"
	CreateSessionRequestUseProxyGeolocationCityLuxembourg              CreateSessionRequestUseProxyGeolocationCity = "LUXEMBOURG"
	CreateSessionRequestUseProxyGeolocationCityLuziania                CreateSessionRequestUseProxyGeolocationCity = "LUZIANIA"
	CreateSessionRequestUseProxyGeolocationCityLviv                    CreateSessionRequestUseProxyGeolocationCity = "LVIV"
	CreateSessionRequestUseProxyGeolocationCityLyon                    CreateSessionRequestUseProxyGeolocationCity = "LYON"
	CreateSessionRequestUseProxyGeolocationCityMabalacat               CreateSessionRequestUseProxyGeolocationCity = "MABALACAT"
	CreateSessionRequestUseProxyGeolocationCityMacae                   CreateSessionRequestUseProxyGeolocationCity = "MACAE"
	CreateSessionRequestUseProxyGeolocationCityMacao                   CreateSessionRequestUseProxyGeolocationCity = "MACAO"
	CreateSessionRequestUseProxyGeolocationCityMacapa                  CreateSessionRequestUseProxyGeolocationCity = "MACAPA"
	CreateSessionRequestUseProxyGeolocationCityMaceio                  CreateSessionRequestUseProxyGeolocationCity = "MACEIO"
	CreateSessionRequestUseProxyGeolocationCityMachala                 CreateSessionRequestUseProxyGeolocationCity = "MACHALA"
	CreateSessionRequestUseProxyGeolocationCityMadison                 CreateSessionRequestUseProxyGeolocationCity = "MADISON"
	CreateSessionRequestUseProxyGeolocationCityMadrid                  CreateSessionRequestUseProxyGeolocationCity = "MADRID"
	CreateSessionRequestUseProxyGeolocationCityMage                    CreateSessionRequestUseProxyGeolocationCity = "MAGE"
	CreateSessionRequestUseProxyGeolocationCityMagelang                CreateSessionRequestUseProxyGeolocationCity = "MAGELANG"
	CreateSessionRequestUseProxyGeolocationCityMagnesiaAdSipylum       CreateSessionRequestUseProxyGeolocationCity = "MAGNESIA_AD SIPYLUM"
	CreateSessionRequestUseProxyGeolocationCityMakassar                CreateSessionRequestUseProxyGeolocationCity = "MAKASSAR"
	CreateSessionRequestUseProxyGeolocationCityMakatiCity              CreateSessionRequestUseProxyGeolocationCity = "MAKATI_CITY"
	CreateSessionRequestUseProxyGeolocationCityMalabon                 CreateSessionRequestUseProxyGeolocationCity = "MALABON"
	CreateSessionRequestUseProxyGeolocationCityMalaga                  CreateSessionRequestUseProxyGeolocationCity = "MALAGA"
	CreateSessionRequestUseProxyGeolocationCityMalang                  CreateSessionRequestUseProxyGeolocationCity = "MALANG"
	CreateSessionRequestUseProxyGeolocationCityMalappuram              CreateSessionRequestUseProxyGeolocationCity = "MALAPPURAM"
	CreateSessionRequestUseProxyGeolocationCityMaldonado               CreateSessionRequestUseProxyGeolocationCity = "MALDONADO"
	CreateSessionRequestUseProxyGeolocationCityMale                    CreateSessionRequestUseProxyGeolocationCity = "MALE"
	CreateSessionRequestUseProxyGeolocationCityMalmo                   CreateSessionRequestUseProxyGeolocationCity = "MALMO"
	CreateSessionRequestUseProxyGeolocationCityManado                  CreateSessionRequestUseProxyGeolocationCity = "MANADO"
	CreateSessionRequestUseProxyGeolocationCityManagua                 CreateSessionRequestUseProxyGeolocationCity = "MANAGUA"
	CreateSessionRequestUseProxyGeolocationCityManama                  CreateSessionRequestUseProxyGeolocationCity = "MANAMA"
	CreateSessionRequestUseProxyGeolocationCityManaus                  CreateSessionRequestUseProxyGeolocationCity = "MANAUS"
	CreateSessionRequestUseProxyGeolocationCityManchester              CreateSessionRequestUseProxyGeolocationCity = "MANCHESTER"
	CreateSessionRequestUseProxyGeolocationCityMandaluyongCity         CreateSessionRequestUseProxyGeolocationCity = "MANDALUYONG_CITY"
	CreateSessionRequestUseProxyGeolocationCityManila                  CreateSessionRequestUseProxyGeolocationCity = "MANILA"
	CreateSessionRequestUseProxyGeolocationCityManizales               CreateSessionRequestUseProxyGeolocationCity = "MANIZALES"
	CreateSessionRequestUseProxyGeolocationCityMannheim                CreateSessionRequestUseProxyGeolocationCity = "MANNHEIM"
	CreateSessionRequestUseProxyGeolocationCityMaputo                  CreateSessionRequestUseProxyGeolocationCity = "MAPUTO"
	CreateSessionRequestUseProxyGeolocationCityMarDelPlata             CreateSessionRequestUseProxyGeolocationCity = "MAR_DEL PLATA"
	CreateSessionRequestUseProxyGeolocationCityMaraba                  CreateSessionRequestUseProxyGeolocationCity = "MARABA"
	CreateSessionRequestUseProxyGeolocationCityMaracaibo               CreateSessionRequestUseProxyGeolocationCity = "MARACAIBO"
	CreateSessionRequestUseProxyGeolocationCityMaracanau               CreateSessionRequestUseProxyGeolocationCity = "MARACANAU"
	CreateSessionRequestUseProxyGeolocationCityMaracay                 CreateSessionRequestUseProxyGeolocationCity = "MARACAY"
	CreateSessionRequestUseProxyGeolocationCityMardin                  CreateSessionRequestUseProxyGeolocationCity = "MARDIN"
	CreateSessionRequestUseProxyGeolocationCityMaribor                 CreateSessionRequestUseProxyGeolocationCity = "MARIBOR"
	CreateSessionRequestUseProxyGeolocationCityMarica                  CreateSessionRequestUseProxyGeolocationCity = "MARICA"
	CreateSessionRequestUseProxyGeolocationCityMarietta                CreateSessionRequestUseProxyGeolocationCity = "MARIETTA"
	CreateSessionRequestUseProxyGeolocationCityMarikinaCity            CreateSessionRequestUseProxyGeolocationCity = "MARIKINA_CITY"
	CreateSessionRequestUseProxyGeolocationCityMarilia                 CreateSessionRequestUseProxyGeolocationCity = "MARILIA"
	CreateSessionRequestUseProxyGeolocationCityMaringa                 CreateSessionRequestUseProxyGeolocationCity = "MARINGA"
	CreateSessionRequestUseProxyGeolocationCityMarrakesh               CreateSessionRequestUseProxyGeolocationCity = "MARRAKESH"
	CreateSessionRequestUseProxyGeolocationCityMarseille               CreateSessionRequestUseProxyGeolocationCity = "MARSEILLE"
	CreateSessionRequestUseProxyGeolocationCityMaua                    CreateSessionRequestUseProxyGeolocationCity = "MAUA"
	CreateSessionRequestUseProxyGeolocationCityMazatlan                CreateSessionRequestUseProxyGeolocationCity = "MAZATLAN"
	CreateSessionRequestUseProxyGeolocationCityMedan                   CreateSessionRequestUseProxyGeolocationCity = "MEDAN"
	CreateSessionRequestUseProxyGeolocationCityMedellin                CreateSessionRequestUseProxyGeolocationCity = "MEDELLIN"
	CreateSessionRequestUseProxyGeolocationCityMedina                  CreateSessionRequestUseProxyGeolocationCity = "MEDINA"
	CreateSessionRequestUseProxyGeolocationCityMeerut                  CreateSessionRequestUseProxyGeolocationCity = "MEERUT"
	CreateSessionRequestUseProxyGeolocationCityMeknes                  CreateSessionRequestUseProxyGeolocationCity = "MEKNES"
	CreateSessionRequestUseProxyGeolocationCityMelbourne               CreateSessionRequestUseProxyGeolocationCity = "MELBOURNE"
	CreateSessionRequestUseProxyGeolocationCityMemphis                 CreateSessionRequestUseProxyGeolocationCity = "MEMPHIS"
	CreateSessionRequestUseProxyGeolocationCityMendoza                 CreateSessionRequestUseProxyGeolocationCity = "MENDOZA"
	CreateSessionRequestUseProxyGeolocationCityMerida                  CreateSessionRequestUseProxyGeolocationCity = "MERIDA"
	CreateSessionRequestUseProxyGeolocationCityMerkez                  CreateSessionRequestUseProxyGeolocationCity = "MERKEZ"
	CreateSessionRequestUseProxyGeolocationCityMerlo                   CreateSessionRequestUseProxyGeolocationCity = "MERLO"
	CreateSessionRequestUseProxyGeolocationCityMersin                  CreateSessionRequestUseProxyGeolocationCity = "MERSIN"
	CreateSessionRequestUseProxyGeolocationCityMesa                    CreateSessionRequestUseProxyGeolocationCity = "MESA"
	CreateSessionRequestUseProxyGeolocationCityMexicali                CreateSessionRequestUseProxyGeolocationCity = "MEXICALI"
	CreateSessionRequestUseProxyGeolocationCityMexicoCity              CreateSessionRequestUseProxyGeolocationCity = "MEXICO_CITY"
	CreateSessionRequestUseProxyGeolocationCityMilan                   CreateSessionRequestUseProxyGeolocationCity = "MILAN"
	CreateSessionRequestUseProxyGeolocationCityMiltonKeynes            CreateSessionRequestUseProxyGeolocationCity = "MILTON_KEYNES"
	CreateSessionRequestUseProxyGeolocationCityMilwaukee               CreateSessionRequestUseProxyGeolocationCity = "MILWAUKEE"
	CreateSessionRequestUseProxyGeolocationCityMinneapolis             CreateSessionRequestUseProxyGeolocationCity = "MINNEAPOLIS"
	CreateSessionRequestUseProxyGeolocationCityMinsk                   CreateSessionRequestUseProxyGeolocationCity = "MINSK"
	CreateSessionRequestUseProxyGeolocationCityMiskolc                 CreateSessionRequestUseProxyGeolocationCity = "MISKOLC"
	CreateSessionRequestUseProxyGeolocationCityMississauga             CreateSessionRequestUseProxyGeolocationCity = "MISSISSAUGA"
	CreateSessionRequestUseProxyGeolocationCityMogiDasCruzes           CreateSessionRequestUseProxyGeolocationCity = "MOGI_DAS CRUZES"
	CreateSessionRequestUseProxyGeolocationCityMohali                  CreateSessionRequestUseProxyGeolocationCity = "MOHALI"
	CreateSessionRequestUseProxyGeolocationCityMonroe                  CreateSessionRequestUseProxyGeolocationCity = "MONROE"
	CreateSessionRequestUseProxyGeolocationCityMonteGrande             CreateSessionRequestUseProxyGeolocationCity = "MONTE_GRANDE"
	CreateSessionRequestUseProxyGeolocationCityMontegoBay              CreateSessionRequestUseProxyGeolocationCity = "MONTEGO_BAY"
	CreateSessionRequestUseProxyGeolocationCityMonterrey               CreateSessionRequestUseProxyGeolocationCity = "MONTERREY"
	CreateSessionRequestUseProxyGeolocationCityMontesClaros            CreateSessionRequestUseProxyGeolocationCity = "MONTES_CLAROS"
	CreateSessionRequestUseProxyGeolocationCityMontevideo              CreateSessionRequestUseProxyGeolocationCity = "MONTEVIDEO"
	CreateSessionRequestUseProxyGeolocationCityMontgomery              CreateSessionRequestUseProxyGeolocationCity = "MONTGOMERY"
	CreateSessionRequestUseProxyGeolocationCityMontpellier             CreateSessionRequestUseProxyGeolocationCity = "MONTPELLIER"
	CreateSessionRequestUseProxyGeolocationCityMontreal                CreateSessionRequestUseProxyGeolocationCity = "MONTREAL"
	CreateSessionRequestUseProxyGeolocationCityMorelia                 CreateSessionRequestUseProxyGeolocationCity = "MORELIA"
	CreateSessionRequestUseProxyGeolocationCityMoreno                  CreateSessionRequestUseProxyGeolocationCity = "MORENO"
	CreateSessionRequestUseProxyGeolocationCityMoron                   CreateSessionRequestUseProxyGeolocationCity = "MORON"
	CreateSessionRequestUseProxyGeolocationCityMossoro                 CreateSessionRequestUseProxyGeolocationCity = "MOSSORO"
	CreateSessionRequestUseProxyGeolocationCityMugla                   CreateSessionRequestUseProxyGeolocationCity = "MUGLA"
	CreateSessionRequestUseProxyGeolocationCityMultan                  CreateSessionRequestUseProxyGeolocationCity = "MULTAN"
	CreateSessionRequestUseProxyGeolocationCityMumbai                  CreateSessionRequestUseProxyGeolocationCity = "MUMBAI"
	CreateSessionRequestUseProxyGeolocationCityMunich                  CreateSessionRequestUseProxyGeolocationCity = "MUNICH"
	CreateSessionRequestUseProxyGeolocationCityMurcia                  CreateSessionRequestUseProxyGeolocationCity = "MURCIA"
	CreateSessionRequestUseProxyGeolocationCityMuscat                  CreateSessionRequestUseProxyGeolocationCity = "MUSCAT"
	CreateSessionRequestUseProxyGeolocationCityMuzaffargarh            CreateSessionRequestUseProxyGeolocationCity = "MUZAFFARGARH"
	CreateSessionRequestUseProxyGeolocationCityMykolayiv               CreateSessionRequestUseProxyGeolocationCity = "MYKOLAYIV"
	CreateSessionRequestUseProxyGeolocationCityNaaldwijk               CreateSessionRequestUseProxyGeolocationCity = "NAALDWIJK"
	CreateSessionRequestUseProxyGeolocationCityNaga                    CreateSessionRequestUseProxyGeolocationCity = "NAGA"
	CreateSessionRequestUseProxyGeolocationCityNagpur                  CreateSessionRequestUseProxyGeolocationCity = "NAGPUR"
	CreateSessionRequestUseProxyGeolocationCityNairobi                 CreateSessionRequestUseProxyGeolocationCity = "NAIROBI"
	CreateSessionRequestUseProxyGeolocationCityNantes                  CreateSessionRequestUseProxyGeolocationCity = "NANTES"
	CreateSessionRequestUseProxyGeolocationCityNaples                  CreateSessionRequestUseProxyGeolocationCity = "NAPLES"
	CreateSessionRequestUseProxyGeolocationCityNashville               CreateSessionRequestUseProxyGeolocationCity = "NASHVILLE"
	CreateSessionRequestUseProxyGeolocationCityNassau                  CreateSessionRequestUseProxyGeolocationCity = "NASSAU"
	CreateSessionRequestUseProxyGeolocationCityNasugbu                 CreateSessionRequestUseProxyGeolocationCity = "NASUGBU"
	CreateSessionRequestUseProxyGeolocationCityNatal                   CreateSessionRequestUseProxyGeolocationCity = "NATAL"
	CreateSessionRequestUseProxyGeolocationCityNaucalpan               CreateSessionRequestUseProxyGeolocationCity = "NAUCALPAN"
	CreateSessionRequestUseProxyGeolocationCityNaviMumbai              CreateSessionRequestUseProxyGeolocationCity = "NAVI_MUMBAI"
	CreateSessionRequestUseProxyGeolocationCityNeiva                   CreateSessionRequestUseProxyGeolocationCity = "NEIVA"
	CreateSessionRequestUseProxyGeolocationCityNeuquen                 CreateSessionRequestUseProxyGeolocationCity = "NEUQUEN"
	CreateSessionRequestUseProxyGeolocationCityNevsehir                CreateSessionRequestUseProxyGeolocationCity = "NEVSEHIR"
	CreateSessionRequestUseProxyGeolocationCityNewDelhi                CreateSessionRequestUseProxyGeolocationCity = "NEW_DELHI"
	CreateSessionRequestUseProxyGeolocationCityNewOrleans              CreateSessionRequestUseProxyGeolocationCity = "NEW_ORLEANS"
	CreateSessionRequestUseProxyGeolocationCityNewTaipei               CreateSessionRequestUseProxyGeolocationCity = "NEW_TAIPEI"
	CreateSessionRequestUseProxyGeolocationCityNewark                  CreateSessionRequestUseProxyGeolocationCity = "NEWARK"
	CreateSessionRequestUseProxyGeolocationCityNewcastleUponTyne       CreateSessionRequestUseProxyGeolocationCity = "NEWCASTLE_UPON TYNE"
	CreateSessionRequestUseProxyGeolocationCityNhaTrang                CreateSessionRequestUseProxyGeolocationCity = "NHA_TRANG"
	CreateSessionRequestUseProxyGeolocationCityNice                    CreateSessionRequestUseProxyGeolocationCity = "NICE"
	CreateSessionRequestUseProxyGeolocationCityNicosia                 CreateSessionRequestUseProxyGeolocationCity = "NICOSIA"
	CreateSessionRequestUseProxyGeolocationCityNilopolis               CreateSessionRequestUseProxyGeolocationCity = "NILOPOLIS"
	CreateSessionRequestUseProxyGeolocationCityNis                     CreateSessionRequestUseProxyGeolocationCity = "NIS"
	CreateSessionRequestUseProxyGeolocationCityNiteroi                 CreateSessionRequestUseProxyGeolocationCity = "NITEROI"
	CreateSessionRequestUseProxyGeolocationCityNitra                   CreateSessionRequestUseProxyGeolocationCity = "NITRA"
	CreateSessionRequestUseProxyGeolocationCityNizhniyNovgorod         CreateSessionRequestUseProxyGeolocationCity = "NIZHNIY_NOVGOROD"
	CreateSessionRequestUseProxyGeolocationCityNogales                 CreateSessionRequestUseProxyGeolocationCity = "NOGALES"
	CreateSessionRequestUseProxyGeolocationCityNoida                   CreateSessionRequestUseProxyGeolocationCity = "NOIDA"
	CreateSessionRequestUseProxyGeolocationCityNorthampton             CreateSessionRequestUseProxyGeolocationCity = "NORTHAMPTON"
	CreateSessionRequestUseProxyGeolocationCityNorwich                 CreateSessionRequestUseProxyGeolocationCity = "NORWICH"
	CreateSessionRequestUseProxyGeolocationCityNottingham              CreateSessionRequestUseProxyGeolocationCity = "NOTTINGHAM"
	CreateSessionRequestUseProxyGeolocationCityNovaFriburgo            CreateSessionRequestUseProxyGeolocationCity = "NOVA_FRIBURGO"
	CreateSessionRequestUseProxyGeolocationCityNovaIguacu              CreateSessionRequestUseProxyGeolocationCity = "NOVA_IGUACU"
	CreateSessionRequestUseProxyGeolocationCityNoviSad                 CreateSessionRequestUseProxyGeolocationCity = "NOVI_SAD"
	CreateSessionRequestUseProxyGeolocationCityNovoHamburgo            CreateSessionRequestUseProxyGeolocationCity = "NOVO_HAMBURGO"
	CreateSessionRequestUseProxyGeolocationCityNovosibirsk             CreateSessionRequestUseProxyGeolocationCity = "NOVOSIBIRSK"
	CreateSessionRequestUseProxyGeolocationCityNuremberg               CreateSessionRequestUseProxyGeolocationCity = "NUREMBERG"
	CreateSessionRequestUseProxyGeolocationCityOakland                 CreateSessionRequestUseProxyGeolocationCity = "OAKLAND"
	CreateSessionRequestUseProxyGeolocationCityOaxacaCity              CreateSessionRequestUseProxyGeolocationCity = "OAXACA_CITY"
	CreateSessionRequestUseProxyGeolocationCityOdesa                   CreateSessionRequestUseProxyGeolocationCity = "ODESA"
	CreateSessionRequestUseProxyGeolocationCityOklahomaCity            CreateSessionRequestUseProxyGeolocationCity = "OKLAHOMA_CITY"
	CreateSessionRequestUseProxyGeolocationCityOlinda                  CreateSessionRequestUseProxyGeolocationCity = "OLINDA"
	CreateSessionRequestUseProxyGeolocationCityOlomouc                 CreateSessionRequestUseProxyGeolocationCity = "OLOMOUC"
	CreateSessionRequestUseProxyGeolocationCityOlongapoCity            CreateSessionRequestUseProxyGeolocationCity = "OLONGAPO_CITY"
	CreateSessionRequestUseProxyGeolocationCityOlsztyn                 CreateSessionRequestUseProxyGeolocationCity = "OLSZTYN"
	CreateSessionRequestUseProxyGeolocationCityOmaha                   CreateSessionRequestUseProxyGeolocationCity = "OMAHA"
	CreateSessionRequestUseProxyGeolocationCityOradea                  CreateSessionRequestUseProxyGeolocationCity = "ORADEA"
	CreateSessionRequestUseProxyGeolocationCityOran                    CreateSessionRequestUseProxyGeolocationCity = "ORAN"
	CreateSessionRequestUseProxyGeolocationCityOrdu                    CreateSessionRequestUseProxyGeolocationCity = "ORDU"
	CreateSessionRequestUseProxyGeolocationCityOrlando                 CreateSessionRequestUseProxyGeolocationCity = "ORLANDO"
	CreateSessionRequestUseProxyGeolocationCityOsasco                  CreateSessionRequestUseProxyGeolocationCity = "OSASCO"
	CreateSessionRequestUseProxyGeolocationCityOslo                    CreateSessionRequestUseProxyGeolocationCity = "OSLO"
	CreateSessionRequestUseProxyGeolocationCityOsmaniye                CreateSessionRequestUseProxyGeolocationCity = "OSMANIYE"
	CreateSessionRequestUseProxyGeolocationCityOstrava                 CreateSessionRequestUseProxyGeolocationCity = "OSTRAVA"
	CreateSessionRequestUseProxyGeolocationCityOttawa                  CreateSessionRequestUseProxyGeolocationCity = "OTTAWA"
	CreateSessionRequestUseProxyGeolocationCityOujda                   CreateSessionRequestUseProxyGeolocationCity = "OUJDA"
	CreateSessionRequestUseProxyGeolocationCityOurinhos                CreateSessionRequestUseProxyGeolocationCity = "OURINHOS"
	CreateSessionRequestUseProxyGeolocationCityPachuca                 CreateSessionRequestUseProxyGeolocationCity = "PACHUCA"
	CreateSessionRequestUseProxyGeolocationCityPadova                  CreateSessionRequestUseProxyGeolocationCity = "PADOVA"
	CreateSessionRequestUseProxyGeolocationCityPalakkad                CreateSessionRequestUseProxyGeolocationCity = "PALAKKAD"
	CreateSessionRequestUseProxyGeolocationCityPalembang               CreateSessionRequestUseProxyGeolocationCity = "PALEMBANG"
	CreateSessionRequestUseProxyGeolocationCityPalermo                 CreateSessionRequestUseProxyGeolocationCity = "PALERMO"
	CreateSessionRequestUseProxyGeolocationCityPalhoca                 CreateSessionRequestUseProxyGeolocationCity = "PALHOCA"
	CreateSessionRequestUseProxyGeolocationCityPalma                   CreateSessionRequestUseProxyGeolocationCity = "PALMA"
	CreateSessionRequestUseProxyGeolocationCityPalmas                  CreateSessionRequestUseProxyGeolocationCity = "PALMAS"
	CreateSessionRequestUseProxyGeolocationCityPanamaCity              CreateSessionRequestUseProxyGeolocationCity = "PANAMA_CITY"
	CreateSessionRequestUseProxyGeolocationCityParamaribo              CreateSessionRequestUseProxyGeolocationCity = "PARAMARIBO"
	CreateSessionRequestUseProxyGeolocationCityParana                  CreateSessionRequestUseProxyGeolocationCity = "PARANA"
	CreateSessionRequestUseProxyGeolocationCityParanagua               CreateSessionRequestUseProxyGeolocationCity = "PARANAGUA"
	CreateSessionRequestUseProxyGeolocationCityParanaqueCity           CreateSessionRequestUseProxyGeolocationCity = "PARANAQUE_CITY"
	CreateSessionRequestUseProxyGeolocationCityParauapebas             CreateSessionRequestUseProxyGeolocationCity = "PARAUAPEBAS"
	CreateSessionRequestUseProxyGeolocationCityParis                   CreateSessionRequestUseProxyGeolocationCity = "PARIS"
	CreateSessionRequestUseProxyGeolocationCityParnaiba                CreateSessionRequestUseProxyGeolocationCity = "PARNAIBA"
	CreateSessionRequestUseProxyGeolocationCityParnamirim              CreateSessionRequestUseProxyGeolocationCity = "PARNAMIRIM"
	CreateSessionRequestUseProxyGeolocationCityPassoFundo              CreateSessionRequestUseProxyGeolocationCity = "PASSO_FUNDO"
	CreateSessionRequestUseProxyGeolocationCityPasto                   CreateSessionRequestUseProxyGeolocationCity = "PASTO"
	CreateSessionRequestUseProxyGeolocationCityPatan                   CreateSessionRequestUseProxyGeolocationCity = "PATAN"
	CreateSessionRequestUseProxyGeolocationCityPatna                   CreateSessionRequestUseProxyGeolocationCity = "PATNA"
	CreateSessionRequestUseProxyGeolocationCityPatosDeMinas            CreateSessionRequestUseProxyGeolocationCity = "PATOS_DE MINAS"
	CreateSessionRequestUseProxyGeolocationCityPaulista                CreateSessionRequestUseProxyGeolocationCity = "PAULISTA"
	CreateSessionRequestUseProxyGeolocationCityPecs                    CreateSessionRequestUseProxyGeolocationCity = "PECS"
	CreateSessionRequestUseProxyGeolocationCityPekanbaru               CreateSessionRequestUseProxyGeolocationCity = "PEKANBARU"
	CreateSessionRequestUseProxyGeolocationCityPelotas                 CreateSessionRequestUseProxyGeolocationCity = "PELOTAS"
	CreateSessionRequestUseProxyGeolocationCityPeoria                  CreateSessionRequestUseProxyGeolocationCity = "PEORIA"
	CreateSessionRequestUseProxyGeolocationCityPereira                 CreateSessionRequestUseProxyGeolocationCity = "PEREIRA"
	CreateSessionRequestUseProxyGeolocationCityPerm                    CreateSessionRequestUseProxyGeolocationCity = "PERM"
	CreateSessionRequestUseProxyGeolocationCityPerth                   CreateSessionRequestUseProxyGeolocationCity = "PERTH"
	CreateSessionRequestUseProxyGeolocationCityPescara                 CreateSessionRequestUseProxyGeolocationCity = "PESCARA"
	CreateSessionRequestUseProxyGeolocationCityPeshawar                CreateSessionRequestUseProxyGeolocationCity = "PESHAWAR"
	CreateSessionRequestUseProxyGeolocationCityPetahTikva              CreateSessionRequestUseProxyGeolocationCity = "PETAH_TIKVA"
	CreateSessionRequestUseProxyGeolocationCityPetalingJaya            CreateSessionRequestUseProxyGeolocationCity = "PETALING_JAYA"
	CreateSessionRequestUseProxyGeolocationCityPetrolina               CreateSessionRequestUseProxyGeolocationCity = "PETROLINA"
	CreateSessionRequestUseProxyGeolocationCityPetropolis              CreateSessionRequestUseProxyGeolocationCity = "PETROPOLIS"
	CreateSessionRequestUseProxyGeolocationCityPhiladelphia            CreateSessionRequestUseProxyGeolocationCity = "PHILADELPHIA"
	CreateSessionRequestUseProxyGeolocationCityPhnomPenh               CreateSessionRequestUseProxyGeolocationCity = "PHNOM_PENH"
	CreateSessionRequestUseProxyGeolocationCityPhoenix                 CreateSessionRequestUseProxyGeolocationCity = "PHOENIX"
	CreateSessionRequestUseProxyGeolocationCityPilar                   CreateSessionRequestUseProxyGeolocationCity = "PILAR"
	CreateSessionRequestUseProxyGeolocationCityPindamonhangaba         CreateSessionRequestUseProxyGeolocationCity = "PINDAMONHANGABA"
	CreateSessionRequestUseProxyGeolocationCityPiracicaba              CreateSessionRequestUseProxyGeolocationCity = "PIRACICABA"
	CreateSessionRequestUseProxyGeolocationCityPitesti                 CreateSessionRequestUseProxyGeolocationCity = "PITESTI"
	CreateSessionRequestUseProxyGeolocationCityPittsburgh              CreateSessionRequestUseProxyGeolocationCity = "PITTSBURGH"
	CreateSessionRequestUseProxyGeolocationCityPiura                   CreateSessionRequestUseProxyGeolocationCity = "PIURA"
	CreateSessionRequestUseProxyGeolocationCityPlano                   CreateSessionRequestUseProxyGeolocationCity = "PLANO"
	CreateSessionRequestUseProxyGeolocationCityPloiesti                CreateSessionRequestUseProxyGeolocationCity = "PLOIESTI"
	CreateSessionRequestUseProxyGeolocationCityPlovdiv                 CreateSessionRequestUseProxyGeolocationCity = "PLOVDIV"
	CreateSessionRequestUseProxyGeolocationCityPlymouth                CreateSessionRequestUseProxyGeolocationCity = "PLYMOUTH"
	CreateSessionRequestUseProxyGeolocationCityPocosDeCaldas           CreateSessionRequestUseProxyGeolocationCity = "POCOS_DE CALDAS"
	CreateSessionRequestUseProxyGeolocationCityPodgorica               CreateSessionRequestUseProxyGeolocationCity = "PODGORICA"
	CreateSessionRequestUseProxyGeolocationCityPoltava                 CreateSessionRequestUseProxyGeolocationCity = "POLTAVA"
	CreateSessionRequestUseProxyGeolocationCityPontaGrossa             CreateSessionRequestUseProxyGeolocationCity = "PONTA_GROSSA"
	CreateSessionRequestUseProxyGeolocationCityPontianak               CreateSessionRequestUseProxyGeolocationCity = "PONTIANAK"
	CreateSessionRequestUseProxyGeolocationCityPopayan                 CreateSessionRequestUseProxyGeolocationCity = "POPAYAN"
	CreateSessionRequestUseProxyGeolocationCityPortAuPrince            CreateSessionRequestUseProxyGeolocationCity = "PORT_AU PRINCE"
	CreateSessionRequestUseProxyGeolocationCityPortElizabeth           CreateSessionRequestUseProxyGeolocationCity = "PORT_ELIZABETH"
	CreateSessionRequestUseProxyGeolocationCityPortHarcourt            CreateSessionRequestUseProxyGeolocationCity = "PORT_HARCOURT"
	CreateSessionRequestUseProxyGeolocationCityPortLouis               CreateSessionRequestUseProxyGeolocationCity = "PORT_LOUIS"
	CreateSessionRequestUseProxyGeolocationCityPortMontt               CreateSessionRequestUseProxyGeolocationCity = "PORT_MONTT"
	CreateSessionRequestUseProxyGeolocationCityPortOfSpain             CreateSessionRequestUseProxyGeolocationCity = "PORT_OF SPAIN"
	CreateSessionRequestUseProxyGeolocationCityPortSaid                CreateSessionRequestUseProxyGeolocationCity = "PORT_SAID"
	CreateSessionRequestUseProxyGeolocationCityPortland                CreateSessionRequestUseProxyGeolocationCity = "PORTLAND"
	CreateSessionRequestUseProxyGeolocationCityPorto                   CreateSessionRequestUseProxyGeolocationCity = "PORTO"
	CreateSessionRequestUseProxyGeolocationCityPortoAlegre             CreateSessionRequestUseProxyGeolocationCity = "PORTO_ALEGRE"
	CreateSessionRequestUseProxyGeolocationCityPortoSeguro             CreateSessionRequestUseProxyGeolocationCity = "PORTO_SEGURO"
	CreateSessionRequestUseProxyGeolocationCityPortoVelho              CreateSessionRequestUseProxyGeolocationCity = "PORTO_VELHO"
	CreateSessionRequestUseProxyGeolocationCityPortoviejo              CreateSessionRequestUseProxyGeolocationCity = "PORTOVIEJO"
	CreateSessionRequestUseProxyGeolocationCityPosadas                 CreateSessionRequestUseProxyGeolocationCity = "POSADAS"
	CreateSessionRequestUseProxyGeolocationCityPousoAlegre             CreateSessionRequestUseProxyGeolocationCity = "POUSO_ALEGRE"
	CreateSessionRequestUseProxyGeolocationCityPoznan                  CreateSessionRequestUseProxyGeolocationCity = "POZNAN"
	CreateSessionRequestUseProxyGeolocationCityPrague                  CreateSessionRequestUseProxyGeolocationCity = "PRAGUE"
	CreateSessionRequestUseProxyGeolocationCityPraiaGrande             CreateSessionRequestUseProxyGeolocationCity = "PRAIA_GRANDE"
	CreateSessionRequestUseProxyGeolocationCityPresidentePrudente      CreateSessionRequestUseProxyGeolocationCity = "PRESIDENTE_PRUDENTE"
	CreateSessionRequestUseProxyGeolocationCityPretoria                CreateSessionRequestUseProxyGeolocationCity = "PRETORIA"
	CreateSessionRequestUseProxyGeolocationCityPristina                CreateSessionRequestUseProxyGeolocationCity = "PRISTINA"
	CreateSessionRequestUseProxyGeolocationCityProvidence              CreateSessionRequestUseProxyGeolocationCity = "PROVIDENCE"
	CreateSessionRequestUseProxyGeolocationCityPucallpa                CreateSessionRequestUseProxyGeolocationCity = "PUCALLPA"
	CreateSessionRequestUseProxyGeolocationCityPuchongBatuDuaBelas     CreateSessionRequestUseProxyGeolocationCity = "PUCHONG_BATU DUA BELAS"
	CreateSessionRequestUseProxyGeolocationCityPueblaCity              CreateSessionRequestUseProxyGeolocationCity = "PUEBLA_CITY"
	CreateSessionRequestUseProxyGeolocationCityPune                    CreateSessionRequestUseProxyGeolocationCity = "PUNE"
	CreateSessionRequestUseProxyGeolocationCityQuebec                  CreateSessionRequestUseProxyGeolocationCity = "QUEBEC"
	CreateSessionRequestUseProxyGeolocationCityQueens                  CreateSessionRequestUseProxyGeolocationCity = "QUEENS"
	CreateSessionRequestUseProxyGeolocationCityQueimados               CreateSessionRequestUseProxyGeolocationCity = "QUEIMADOS"
	CreateSessionRequestUseProxyGeolocationCityQueretaroCity           CreateSessionRequestUseProxyGeolocationCity = "QUERETARO_CITY"
	CreateSessionRequestUseProxyGeolocationCityQuezonCity              CreateSessionRequestUseProxyGeolocationCity = "QUEZON_CITY"
	CreateSessionRequestUseProxyGeolocationCityQuilmes                 CreateSessionRequestUseProxyGeolocationCity = "QUILMES"
	CreateSessionRequestUseProxyGeolocationCityQuito                   CreateSessionRequestUseProxyGeolocationCity = "QUITO"
	CreateSessionRequestUseProxyGeolocationCityRabat                   CreateSessionRequestUseProxyGeolocationCity = "RABAT"
	CreateSessionRequestUseProxyGeolocationCityRaipur                  CreateSessionRequestUseProxyGeolocationCity = "RAIPUR"
	CreateSessionRequestUseProxyGeolocationCityRajkot                  CreateSessionRequestUseProxyGeolocationCity = "RAJKOT"
	CreateSessionRequestUseProxyGeolocationCityRajshahi                CreateSessionRequestUseProxyGeolocationCity = "RAJSHAHI"
	CreateSessionRequestUseProxyGeolocationCityRaleigh                 CreateSessionRequestUseProxyGeolocationCity = "RALEIGH"
	CreateSessionRequestUseProxyGeolocationCityRamatGan                CreateSessionRequestUseProxyGeolocationCity = "RAMAT_GAN"
	CreateSessionRequestUseProxyGeolocationCityRancagua                CreateSessionRequestUseProxyGeolocationCity = "RANCAGUA"
	CreateSessionRequestUseProxyGeolocationCityRanchi                  CreateSessionRequestUseProxyGeolocationCity = "RANCHI"
	CreateSessionRequestUseProxyGeolocationCityRasAlKhaimah            CreateSessionRequestUseProxyGeolocationCity = "RAS_AL KHAIMAH"
	CreateSessionRequestUseProxyGeolocationCityRawalpindi              CreateSessionRequestUseProxyGeolocationCity = "RAWALPINDI"
	CreateSessionRequestUseProxyGeolocationCityReading                 CreateSessionRequestUseProxyGeolocationCity = "READING"
	CreateSessionRequestUseProxyGeolocationCityRecife                  CreateSessionRequestUseProxyGeolocationCity = "RECIFE"
	CreateSessionRequestUseProxyGeolocationCityRegina                  CreateSessionRequestUseProxyGeolocationCity = "REGINA"
	CreateSessionRequestUseProxyGeolocationCityRennes                  CreateSessionRequestUseProxyGeolocationCity = "RENNES"
	CreateSessionRequestUseProxyGeolocationCityReno                    CreateSessionRequestUseProxyGeolocationCity = "RENO"
	CreateSessionRequestUseProxyGeolocationCityResistencia             CreateSessionRequestUseProxyGeolocationCity = "RESISTENCIA"
	CreateSessionRequestUseProxyGeolocationCityReykjavik               CreateSessionRequestUseProxyGeolocationCity = "REYKJAVIK"
	CreateSessionRequestUseProxyGeolocationCityReynosa                 CreateSessionRequestUseProxyGeolocationCity = "REYNOSA"
	CreateSessionRequestUseProxyGeolocationCityRibeiraoDasNeves        CreateSessionRequestUseProxyGeolocationCity = "RIBEIRAO_DAS NEVES"
	CreateSessionRequestUseProxyGeolocationCityRibeiraoPreto           CreateSessionRequestUseProxyGeolocationCity = "RIBEIRAO_PRETO"
	CreateSessionRequestUseProxyGeolocationCityRichmond                CreateSessionRequestUseProxyGeolocationCity = "RICHMOND"
	CreateSessionRequestUseProxyGeolocationCityRiga                    CreateSessionRequestUseProxyGeolocationCity = "RIGA"
	CreateSessionRequestUseProxyGeolocationCityRioBranco               CreateSessionRequestUseProxyGeolocationCity = "RIO_BRANCO"
	CreateSessionRequestUseProxyGeolocationCityRioClaro                CreateSessionRequestUseProxyGeolocationCity = "RIO_CLARO"
	CreateSessionRequestUseProxyGeolocationCityRioCuarto               CreateSessionRequestUseProxyGeolocationCity = "RIO_CUARTO"
	CreateSessionRequestUseProxyGeolocationCityRioDeJaneiro            CreateSessionRequestUseProxyGeolocationCity = "RIO_DE JANEIRO"
	CreateSessionRequestUseProxyGeolocationCityRioDoSul                CreateSessionRequestUseProxyGeolocationCity = "RIO_DO SUL"
	CreateSessionRequestUseProxyGeolocationCityRioGallegos             CreateSessionRequestUseProxyGeolocationCity = "RIO_GALLEGOS"
	CreateSessionRequestUseProxyGeolocationCityRioGrande               CreateSessionRequestUseProxyGeolocationCity = "RIO_GRANDE"
	CreateSessionRequestUseProxyGeolocationCityRishonLetsiyyon         CreateSessionRequestUseProxyGeolocationCity = "RISHON_LETSIYYON"
	CreateSessionRequestUseProxyGeolocationCityRiverside               CreateSessionRequestUseProxyGeolocationCity = "RIVERSIDE"
	CreateSessionRequestUseProxyGeolocationCityRiyadh                  CreateSessionRequestUseProxyGeolocationCity = "RIYADH"
	CreateSessionRequestUseProxyGeolocationCityRize                    CreateSessionRequestUseProxyGeolocationCity = "RIZE"
	CreateSessionRequestUseProxyGeolocationCityRochester               CreateSessionRequestUseProxyGeolocationCity = "ROCHESTER"
	CreateSessionRequestUseProxyGeolocationCityRome                    CreateSessionRequestUseProxyGeolocationCity = "ROME"
	CreateSessionRequestUseProxyGeolocationCityRondonopolis            CreateSessionRequestUseProxyGeolocationCity = "RONDONOPOLIS"
	CreateSessionRequestUseProxyGeolocationCityRosario                 CreateSessionRequestUseProxyGeolocationCity = "ROSARIO"
	CreateSessionRequestUseProxyGeolocationCityRoseau                  CreateSessionRequestUseProxyGeolocationCity = "ROSEAU"
	CreateSessionRequestUseProxyGeolocationCityRostovOnDon             CreateSessionRequestUseProxyGeolocationCity = "ROSTOV_ON DON"
	CreateSessionRequestUseProxyGeolocationCityRotterdam               CreateSessionRequestUseProxyGeolocationCity = "ROTTERDAM"
	CreateSessionRequestUseProxyGeolocationCityRouen                   CreateSessionRequestUseProxyGeolocationCity = "ROUEN"
	CreateSessionRequestUseProxyGeolocationCityRousse                  CreateSessionRequestUseProxyGeolocationCity = "ROUSSE"
	CreateSessionRequestUseProxyGeolocationCityRzeszow                 CreateSessionRequestUseProxyGeolocationCity = "RZESZOW"
	CreateSessionRequestUseProxyGeolocationCitySacramento              CreateSessionRequestUseProxyGeolocationCity = "SACRAMENTO"
	CreateSessionRequestUseProxyGeolocationCitySagar                   CreateSessionRequestUseProxyGeolocationCity = "SAGAR"
	CreateSessionRequestUseProxyGeolocationCitySaintPaul               CreateSessionRequestUseProxyGeolocationCity = "SAINT_PAUL"
	CreateSessionRequestUseProxyGeolocationCitySale                    CreateSessionRequestUseProxyGeolocationCity = "SALE"
	CreateSessionRequestUseProxyGeolocationCitySaltLakeCity            CreateSessionRequestUseProxyGeolocationCity = "SALT_LAKE CITY"
	CreateSessionRequestUseProxyGeolocationCitySalta                   CreateSessionRequestUseProxyGeolocationCity = "SALTA"
	CreateSessionRequestUseProxyGeolocationCitySaltillo                CreateSessionRequestUseProxyGeolocationCity = "SALTILLO"
	CreateSessionRequestUseProxyGeolocationCitySalvador                CreateSessionRequestUseProxyGeolocationCity = "SALVADOR"
	CreateSessionRequestUseProxyGeolocationCitySamara                  CreateSessionRequestUseProxyGeolocationCity = "SAMARA"
	CreateSessionRequestUseProxyGeolocationCitySamarinda               CreateSessionRequestUseProxyGeolocationCity = "SAMARINDA"
	CreateSessionRequestUseProxyGeolocationCitySamarkand               CreateSessionRequestUseProxyGeolocationCity = "SAMARKAND"
	CreateSessionRequestUseProxyGeolocationCitySamsun                  CreateSessionRequestUseProxyGeolocationCity = "SAMSUN"
	CreateSessionRequestUseProxyGeolocationCitySanAntonio              CreateSessionRequestUseProxyGeolocationCity = "SAN_ANTONIO"
	CreateSessionRequestUseProxyGeolocationCitySanDiego                CreateSessionRequestUseProxyGeolocationCity = "SAN_DIEGO"
	CreateSessionRequestUseProxyGeolocationCitySanFernando             CreateSessionRequestUseProxyGeolocationCity = "SAN_FERNANDO"
	CreateSessionRequestUseProxyGeolocationCitySanFrancisco            CreateSessionRequestUseProxyGeolocationCity = "SAN_FRANCISCO"
	CreateSessionRequestUseProxyGeolocationCitySanJose                 CreateSessionRequestUseProxyGeolocationCity = "SAN_JOSE"
	CreateSessionRequestUseProxyGeolocationCitySanJoseDelMonte         CreateSessionRequestUseProxyGeolocationCity = "SAN_JOSE DEL MONTE"
	CreateSessionRequestUseProxyGeolocationCitySanJuan                 CreateSessionRequestUseProxyGeolocationCity = "SAN_JUAN"
	CreateSessionRequestUseProxyGeolocationCitySanJusto                CreateSessionRequestUseProxyGeolocationCity = "SAN_JUSTO"
	CreateSessionRequestUseProxyGeolocationCitySanLuis                 CreateSessionRequestUseProxyGeolocationCity = "SAN_LUIS"
	CreateSessionRequestUseProxyGeolocationCitySanLuisPotosiCity       CreateSessionRequestUseProxyGeolocationCity = "SAN_LUIS POTOSI CITY"
	CreateSessionRequestUseProxyGeolocationCitySanMiguel               CreateSessionRequestUseProxyGeolocationCity = "SAN_MIGUEL"
	CreateSessionRequestUseProxyGeolocationCitySanMiguelDeTucuman      CreateSessionRequestUseProxyGeolocationCity = "SAN_MIGUEL DE TUCUMAN"
	CreateSessionRequestUseProxyGeolocationCitySanPabloCity            CreateSessionRequestUseProxyGeolocationCity = "SAN_PABLO CITY"
	CreateSessionRequestUseProxyGeolocationCitySanPedro                CreateSessionRequestUseProxyGeolocationCity = "SAN_PEDRO"
	CreateSessionRequestUseProxyGeolocationCitySanPedroSula            CreateSessionRequestUseProxyGeolocationCity = "SAN_PEDRO SULA"
	CreateSessionRequestUseProxyGeolocationCitySanSalvador             CreateSessionRequestUseProxyGeolocationCity = "SAN_SALVADOR"
	CreateSessionRequestUseProxyGeolocationCitySanSalvadorDeJujuy      CreateSessionRequestUseProxyGeolocationCity = "SAN_SALVADOR DE JUJUY"
	CreateSessionRequestUseProxyGeolocationCitySanaa                   CreateSessionRequestUseProxyGeolocationCity = "SANAA"
	CreateSessionRequestUseProxyGeolocationCitySanliurfa               CreateSessionRequestUseProxyGeolocationCity = "SANLIURFA"
	CreateSessionRequestUseProxyGeolocationCitySantaCruz               CreateSessionRequestUseProxyGeolocationCity = "SANTA_CRUZ"
	CreateSessionRequestUseProxyGeolocationCitySantaCruzDeTenerife     CreateSessionRequestUseProxyGeolocationCity = "SANTA_CRUZ DE TENERIFE"
	CreateSessionRequestUseProxyGeolocationCitySantaCruzDoSul          CreateSessionRequestUseProxyGeolocationCity = "SANTA_CRUZ DO SUL"
	CreateSessionRequestUseProxyGeolocationCitySantaFe                 CreateSessionRequestUseProxyGeolocationCity = "SANTA_FE"
	CreateSessionRequestUseProxyGeolocationCitySantaLuzia              CreateSessionRequestUseProxyGeolocationCity = "SANTA_LUZIA"
	CreateSessionRequestUseProxyGeolocationCitySantaMaria              CreateSessionRequestUseProxyGeolocationCity = "SANTA_MARIA"
	CreateSessionRequestUseProxyGeolocationCitySantaMarta              CreateSessionRequestUseProxyGeolocationCity = "SANTA_MARTA"
	CreateSessionRequestUseProxyGeolocationCitySantaRosa               CreateSessionRequestUseProxyGeolocationCity = "SANTA_ROSA"
	CreateSessionRequestUseProxyGeolocationCitySantarem                CreateSessionRequestUseProxyGeolocationCity = "SANTAREM"
	CreateSessionRequestUseProxyGeolocationCitySantiago                CreateSessionRequestUseProxyGeolocationCity = "SANTIAGO"
	CreateSessionRequestUseProxyGeolocationCitySantiagoDeCali          CreateSessionRequestUseProxyGeolocationCity = "SANTIAGO_DE CALI"
	CreateSessionRequestUseProxyGeolocationCitySantiagoDeLosCaballeros CreateSessionRequestUseProxyGeolocationCity = "SANTIAGO_DE LOS CABALLEROS"
	CreateSessionRequestUseProxyGeolocationCitySantoAndre              CreateSessionRequestUseProxyGeolocationCity = "SANTO_ANDRE"
	CreateSessionRequestUseProxyGeolocationCitySantoDomingo            CreateSessionRequestUseProxyGeolocationCity = "SANTO_DOMINGO"
	CreateSessionRequestUseProxyGeolocationCitySantoDomingoEste        CreateSessionRequestUseProxyGeolocationCity = "SANTO_DOMINGO ESTE"
	CreateSessionRequestUseProxyGeolocationCitySantos                  CreateSessionRequestUseProxyGeolocationCity = "SANTOS"
	CreateSessionRequestUseProxyGeolocationCitySaoBernardoDoCampo      CreateSessionRequestUseProxyGeolocationCity = "SAO_BERNARDO DO CAMPO"
	CreateSessionRequestUseProxyGeolocationCitySaoCarlos               CreateSessionRequestUseProxyGeolocationCity = "SAO_CARLOS"
	CreateSessionRequestUseProxyGeolocationCitySaoGoncalo              CreateSessionRequestUseProxyGeolocationCity = "SAO_GONCALO"
	CreateSessionRequestUseProxyGeolocationCitySaoJoaoDeMeriti         CreateSessionRequestUseProxyGeolocationCity = "SAO_JOAO DE MERITI"
	CreateSessionRequestUseProxyGeolocationCitySaoJose                 CreateSessionRequestUseProxyGeolocationCity = "SAO_JOSE"
	CreateSessionRequestUseProxyGeolocationCitySaoJoseDoRioPreto       CreateSessionRequestUseProxyGeolocationCity = "SAO_JOSE DO RIO PRETO"
	CreateSessionRequestUseProxyGeolocationCitySaoJoseDosCampos        CreateSessionRequestUseProxyGeolocationCity = "SAO_JOSE DOS CAMPOS"
	CreateSessionRequestUseProxyGeolocationCitySaoJoseDosPinhais       CreateSessionRequestUseProxyGeolocationCity = "SAO_JOSE DOS PINHAIS"
	CreateSessionRequestUseProxyGeolocationCitySaoLeopoldo             CreateSessionRequestUseProxyGeolocationCity = "SAO_LEOPOLDO"
	CreateSessionRequestUseProxyGeolocationCitySaoLuis                 CreateSessionRequestUseProxyGeolocationCity = "SAO_LUIS"
	CreateSessionRequestUseProxyGeolocationCitySaoPaulo                CreateSessionRequestUseProxyGeolocationCity = "SAO_PAULO"
	CreateSessionRequestUseProxyGeolocationCitySaoVicente              CreateSessionRequestUseProxyGeolocationCity = "SAO_VICENTE"
	CreateSessionRequestUseProxyGeolocationCitySarajevo                CreateSessionRequestUseProxyGeolocationCity = "SARAJEVO"
	CreateSessionRequestUseProxyGeolocationCitySaskatoon               CreateSessionRequestUseProxyGeolocationCity = "SASKATOON"
	CreateSessionRequestUseProxyGeolocationCityScarborough             CreateSessionRequestUseProxyGeolocationCity = "SCARBOROUGH"
	CreateSessionRequestUseProxyGeolocationCitySeattle                 CreateSessionRequestUseProxyGeolocationCity = "SEATTLE"
	CreateSessionRequestUseProxyGeolocationCitySemarang                CreateSessionRequestUseProxyGeolocationCity = "SEMARANG"
	CreateSessionRequestUseProxyGeolocationCitySeoGu                   CreateSessionRequestUseProxyGeolocationCity = "SEO_GU"
	CreateSessionRequestUseProxyGeolocationCitySeongnamSi              CreateSessionRequestUseProxyGeolocationCity = "SEONGNAM_SI"
	CreateSessionRequestUseProxyGeolocationCitySeoul                   CreateSessionRequestUseProxyGeolocationCity = "SEOUL"
	CreateSessionRequestUseProxyGeolocationCitySerra                   CreateSessionRequestUseProxyGeolocationCity = "SERRA"
	CreateSessionRequestUseProxyGeolocationCitySeteLagoas              CreateSessionRequestUseProxyGeolocationCity = "SETE_LAGOAS"
	CreateSessionRequestUseProxyGeolocationCitySetif                   CreateSessionRequestUseProxyGeolocationCity = "SETIF"
	CreateSessionRequestUseProxyGeolocationCitySetubal                 CreateSessionRequestUseProxyGeolocationCity = "SETUBAL"
	CreateSessionRequestUseProxyGeolocationCitySeville                 CreateSessionRequestUseProxyGeolocationCity = "SEVILLE"
	CreateSessionRequestUseProxyGeolocationCitySfax                    CreateSessionRequestUseProxyGeolocationCity = "SFAX"
	CreateSessionRequestUseProxyGeolocationCityShahAlam                CreateSessionRequestUseProxyGeolocationCity = "SHAH_ALAM"
	CreateSessionRequestUseProxyGeolocationCityShanghai                CreateSessionRequestUseProxyGeolocationCity = "SHANGHAI"
	CreateSessionRequestUseProxyGeolocationCitySharjah                 CreateSessionRequestUseProxyGeolocationCity = "SHARJAH"
	CreateSessionRequestUseProxyGeolocationCitySheffield               CreateSessionRequestUseProxyGeolocationCity = "SHEFFIELD"
	CreateSessionRequestUseProxyGeolocationCityShenzhen                CreateSessionRequestUseProxyGeolocationCity = "SHENZHEN"
	CreateSessionRequestUseProxyGeolocationCityShimla                  CreateSessionRequestUseProxyGeolocationCity = "SHIMLA"
	CreateSessionRequestUseProxyGeolocationCitySiauliai                CreateSessionRequestUseProxyGeolocationCity = "SIAULIAI"
	CreateSessionRequestUseProxyGeolocationCitySibiu                   CreateSessionRequestUseProxyGeolocationCity = "SIBIU"
	CreateSessionRequestUseProxyGeolocationCitySidoarjo                CreateSessionRequestUseProxyGeolocationCity = "SIDOARJO"
	CreateSessionRequestUseProxyGeolocationCitySikar                   CreateSessionRequestUseProxyGeolocationCity = "SIKAR"
	CreateSessionRequestUseProxyGeolocationCitySilverSpring            CreateSessionRequestUseProxyGeolocationCity = "SILVER_SPRING"
	CreateSessionRequestUseProxyGeolocationCitySinop                   CreateSessionRequestUseProxyGeolocationCity = "SINOP"
	CreateSessionRequestUseProxyGeolocationCitySivas                   CreateSessionRequestUseProxyGeolocationCity = "SIVAS"
	CreateSessionRequestUseProxyGeolocationCitySkikda                  CreateSessionRequestUseProxyGeolocationCity = "SKIKDA"
	CreateSessionRequestUseProxyGeolocationCitySkopje                  CreateSessionRequestUseProxyGeolocationCity = "SKOPJE"
	CreateSessionRequestUseProxyGeolocationCitySlough                  CreateSessionRequestUseProxyGeolocationCity = "SLOUGH"
	CreateSessionRequestUseProxyGeolocationCitySobral                  CreateSessionRequestUseProxyGeolocationCity = "SOBRAL"
	CreateSessionRequestUseProxyGeolocationCitySofia                   CreateSessionRequestUseProxyGeolocationCity = "SOFIA"
	CreateSessionRequestUseProxyGeolocationCitySorocaba                CreateSessionRequestUseProxyGeolocationCity = "SOROCABA"
	CreateSessionRequestUseProxyGeolocationCitySousse                  CreateSessionRequestUseProxyGeolocationCity = "SOUSSE"
	CreateSessionRequestUseProxyGeolocationCitySouthTangerang          CreateSessionRequestUseProxyGeolocationCity = "SOUTH_TANGERANG"
	CreateSessionRequestUseProxyGeolocationCitySouthampton             CreateSessionRequestUseProxyGeolocationCity = "SOUTHAMPTON"
	CreateSessionRequestUseProxyGeolocationCitySouthwark               CreateSessionRequestUseProxyGeolocationCity = "SOUTHWARK"
	CreateSessionRequestUseProxyGeolocationCitySplit                   CreateSessionRequestUseProxyGeolocationCity = "SPLIT"
	CreateSessionRequestUseProxyGeolocationCitySpokane                 CreateSessionRequestUseProxyGeolocationCity = "SPOKANE"
	CreateSessionRequestUseProxyGeolocationCitySpring                  CreateSessionRequestUseProxyGeolocationCity = "SPRING"
	CreateSessionRequestUseProxyGeolocationCitySpringfield             CreateSessionRequestUseProxyGeolocationCity = "SPRINGFIELD"
	CreateSessionRequestUseProxyGeolocationCityStLouis                 CreateSessionRequestUseProxyGeolocationCity = "ST_LOUIS"
	CreateSessionRequestUseProxyGeolocationCityStPetersburg            CreateSessionRequestUseProxyGeolocationCity = "ST_PETERSBURG"
	CreateSessionRequestUseProxyGeolocationCityStaraZagora             CreateSessionRequestUseProxyGeolocationCity = "STARA_ZAGORA"
	CreateSessionRequestUseProxyGeolocationCityStatenIsland            CreateSessionRequestUseProxyGeolocationCity = "STATEN_ISLAND"
	CreateSessionRequestUseProxyGeolocationCityStockholm               CreateSessionRequestUseProxyGeolocationCity = "STOCKHOLM"
	CreateSessionRequestUseProxyGeolocationCityStockton                CreateSessionRequestUseProxyGeolocationCity = "STOCKTON"
	CreateSessionRequestUseProxyGeolocationCityStokeOnTrent            CreateSessionRequestUseProxyGeolocationCity = "STOKE_ON TRENT"
	CreateSessionRequestUseProxyGeolocationCityStrasbourg              CreateSessionRequestUseProxyGeolocationCity = "STRASBOURG"
	CreateSessionRequestUseProxyGeolocationCityStuttgart               CreateSessionRequestUseProxyGeolocationCity = "STUTTGART"
	CreateSessionRequestUseProxyGeolocationCitySumare                  CreateSessionRequestUseProxyGeolocationCity = "SUMARE"
	CreateSessionRequestUseProxyGeolocationCitySurabaya                CreateSessionRequestUseProxyGeolocationCity = "SURABAYA"
	CreateSessionRequestUseProxyGeolocationCitySurakarta               CreateSessionRequestUseProxyGeolocationCity = "SURAKARTA"
	CreateSessionRequestUseProxyGeolocationCitySurat                   CreateSessionRequestUseProxyGeolocationCity = "SURAT"
	CreateSessionRequestUseProxyGeolocationCitySurrey                  CreateSessionRequestUseProxyGeolocationCity = "SURREY"
	CreateSessionRequestUseProxyGeolocationCitySuwon                   CreateSessionRequestUseProxyGeolocationCity = "SUWON"
	CreateSessionRequestUseProxyGeolocationCitySuzano                  CreateSessionRequestUseProxyGeolocationCity = "SUZANO"
	CreateSessionRequestUseProxyGeolocationCitySydney                  CreateSessionRequestUseProxyGeolocationCity = "SYDNEY"
	CreateSessionRequestUseProxyGeolocationCitySzczecin                CreateSessionRequestUseProxyGeolocationCity = "SZCZECIN"
	CreateSessionRequestUseProxyGeolocationCitySzeged                  CreateSessionRequestUseProxyGeolocationCity = "SZEGED"
	CreateSessionRequestUseProxyGeolocationCitySzekesfehervar          CreateSessionRequestUseProxyGeolocationCity = "SZEKESFEHERVAR"
	CreateSessionRequestUseProxyGeolocationCityTaboaoDaSerra           CreateSessionRequestUseProxyGeolocationCity = "TABOAO_DA SERRA"
	CreateSessionRequestUseProxyGeolocationCityTacna                   CreateSessionRequestUseProxyGeolocationCity = "TACNA"
	CreateSessionRequestUseProxyGeolocationCityTacoma                  CreateSessionRequestUseProxyGeolocationCity = "TACOMA"
	CreateSessionRequestUseProxyGeolocationCityTaguig                  CreateSessionRequestUseProxyGeolocationCity = "TAGUIG"
	CreateSessionRequestUseProxyGeolocationCityTaichung                CreateSessionRequestUseProxyGeolocationCity = "TAICHUNG"
	CreateSessionRequestUseProxyGeolocationCityTainanCity              CreateSessionRequestUseProxyGeolocationCity = "TAINAN_CITY"
	CreateSessionRequestUseProxyGeolocationCityTaipei                  CreateSessionRequestUseProxyGeolocationCity = "TAIPEI"
	CreateSessionRequestUseProxyGeolocationCityTalavera                CreateSessionRequestUseProxyGeolocationCity = "TALAVERA"
	CreateSessionRequestUseProxyGeolocationCityTalca                   CreateSessionRequestUseProxyGeolocationCity = "TALCA"
	CreateSessionRequestUseProxyGeolocationCityTallahassee             CreateSessionRequestUseProxyGeolocationCity = "TALLAHASSEE"
	CreateSessionRequestUseProxyGeolocationCityTallinn                 CreateSessionRequestUseProxyGeolocationCity = "TALLINN"
	CreateSessionRequestUseProxyGeolocationCityTampa                   CreateSessionRequestUseProxyGeolocationCity = "TAMPA"
	CreateSessionRequestUseProxyGeolocationCityTampere                 CreateSessionRequestUseProxyGeolocationCity = "TAMPERE"
	CreateSessionRequestUseProxyGeolocationCityTampico                 CreateSessionRequestUseProxyGeolocationCity = "TAMPICO"
	CreateSessionRequestUseProxyGeolocationCityTangerang               CreateSessionRequestUseProxyGeolocationCity = "TANGERANG"
	CreateSessionRequestUseProxyGeolocationCityTangier                 CreateSessionRequestUseProxyGeolocationCity = "TANGIER"
	CreateSessionRequestUseProxyGeolocationCityTanta                   CreateSessionRequestUseProxyGeolocationCity = "TANTA"
	CreateSessionRequestUseProxyGeolocationCityTanza                   CreateSessionRequestUseProxyGeolocationCity = "TANZA"
	CreateSessionRequestUseProxyGeolocationCityTaoyuanDistrict         CreateSessionRequestUseProxyGeolocationCity = "TAOYUAN_DISTRICT"
	CreateSessionRequestUseProxyGeolocationCityTappahannock            CreateSessionRequestUseProxyGeolocationCity = "TAPPAHANNOCK"
	CreateSessionRequestUseProxyGeolocationCityTarlacCity              CreateSessionRequestUseProxyGeolocationCity = "TARLAC_CITY"
	CreateSessionRequestUseProxyGeolocationCityTashkent                CreateSessionRequestUseProxyGeolocationCity = "TASHKENT"
	CreateSessionRequestUseProxyGeolocationCityTasikmalaya             CreateSessionRequestUseProxyGeolocationCity = "TASIKMALAYA"
	CreateSessionRequestUseProxyGeolocationCityTatui                   CreateSessionRequestUseProxyGeolocationCity = "TATUI"
	CreateSessionRequestUseProxyGeolocationCityTaubate                 CreateSessionRequestUseProxyGeolocationCity = "TAUBATE"
	CreateSessionRequestUseProxyGeolocationCityTbilisi                 CreateSessionRequestUseProxyGeolocationCity = "TBILISI"
	CreateSessionRequestUseProxyGeolocationCityTegucigalpa             CreateSessionRequestUseProxyGeolocationCity = "TEGUCIGALPA"
	CreateSessionRequestUseProxyGeolocationCityTehran                  CreateSessionRequestUseProxyGeolocationCity = "TEHRAN"
	CreateSessionRequestUseProxyGeolocationCityTeixeiraDeFreitas       CreateSessionRequestUseProxyGeolocationCity = "TEIXEIRA_DE FREITAS"
	CreateSessionRequestUseProxyGeolocationCityTekirdag                CreateSessionRequestUseProxyGeolocationCity = "TEKIRDAG"
	CreateSessionRequestUseProxyGeolocationCityTelAviv                 CreateSessionRequestUseProxyGeolocationCity = "TEL_AVIV"
	CreateSessionRequestUseProxyGeolocationCityTemuco                  CreateSessionRequestUseProxyGeolocationCity = "TEMUCO"
	CreateSessionRequestUseProxyGeolocationCityTepic                   CreateSessionRequestUseProxyGeolocationCity = "TEPIC"
	CreateSessionRequestUseProxyGeolocationCityTeresina                CreateSessionRequestUseProxyGeolocationCity = "TERESINA"
	CreateSessionRequestUseProxyGeolocationCityTernopil                CreateSessionRequestUseProxyGeolocationCity = "TERNOPIL"
	CreateSessionRequestUseProxyGeolocationCityTerrassa                CreateSessionRequestUseProxyGeolocationCity = "TERRASSA"
	CreateSessionRequestUseProxyGeolocationCityTetouan                 CreateSessionRequestUseProxyGeolocationCity = "TETOUAN"
	CreateSessionRequestUseProxyGeolocationCityThane                   CreateSessionRequestUseProxyGeolocationCity = "THANE"
	CreateSessionRequestUseProxyGeolocationCityTheBronx                CreateSessionRequestUseProxyGeolocationCity = "THE_BRONX"
	CreateSessionRequestUseProxyGeolocationCityTheHague                CreateSessionRequestUseProxyGeolocationCity = "THE_HAGUE"
	CreateSessionRequestUseProxyGeolocationCityThessaloniki            CreateSessionRequestUseProxyGeolocationCity = "THESSALONIKI"
	CreateSessionRequestUseProxyGeolocationCityThiruvananthapuram      CreateSessionRequestUseProxyGeolocationCity = "THIRUVANANTHAPURAM"
	CreateSessionRequestUseProxyGeolocationCityThrissur                CreateSessionRequestUseProxyGeolocationCity = "THRISSUR"
	CreateSessionRequestUseProxyGeolocationCityTijuana                 CreateSessionRequestUseProxyGeolocationCity = "TIJUANA"
	CreateSessionRequestUseProxyGeolocationCityTimisoara               CreateSessionRequestUseProxyGeolocationCity = "TIMISOARA"
	CreateSessionRequestUseProxyGeolocationCityTirana                  CreateSessionRequestUseProxyGeolocationCity = "TIRANA"
	CreateSessionRequestUseProxyGeolocationCityTlalnepantla            CreateSessionRequestUseProxyGeolocationCity = "TLALNEPANTLA"
	CreateSessionRequestUseProxyGeolocationCityTlaxcalaCity            CreateSessionRequestUseProxyGeolocationCity = "TLAXCALA_CITY"
	CreateSessionRequestUseProxyGeolocationCityTlemcen                 CreateSessionRequestUseProxyGeolocationCity = "TLEMCEN"
	CreateSessionRequestUseProxyGeolocationCityTokatProvince           CreateSessionRequestUseProxyGeolocationCity = "TOKAT_PROVINCE"
	CreateSessionRequestUseProxyGeolocationCityTokyo                   CreateSessionRequestUseProxyGeolocationCity = "TOKYO"
	CreateSessionRequestUseProxyGeolocationCityToluca                  CreateSessionRequestUseProxyGeolocationCity = "TOLUCA"
	CreateSessionRequestUseProxyGeolocationCityToronto                 CreateSessionRequestUseProxyGeolocationCity = "TORONTO"
	CreateSessionRequestUseProxyGeolocationCityTorreon                 CreateSessionRequestUseProxyGeolocationCity = "TORREON"
	CreateSessionRequestUseProxyGeolocationCityToulouse                CreateSessionRequestUseProxyGeolocationCity = "TOULOUSE"
	CreateSessionRequestUseProxyGeolocationCityTrabzon                 CreateSessionRequestUseProxyGeolocationCity = "TRABZON"
	CreateSessionRequestUseProxyGeolocationCityTrujillo                CreateSessionRequestUseProxyGeolocationCity = "TRUJILLO"
	CreateSessionRequestUseProxyGeolocationCityTubarao                 CreateSessionRequestUseProxyGeolocationCity = "TUBARAO"
	CreateSessionRequestUseProxyGeolocationCityTucson                  CreateSessionRequestUseProxyGeolocationCity = "TUCSON"
	CreateSessionRequestUseProxyGeolocationCityTuguegaraoCity          CreateSessionRequestUseProxyGeolocationCity = "TUGUEGARAO_CITY"
	CreateSessionRequestUseProxyGeolocationCityTulsa                   CreateSessionRequestUseProxyGeolocationCity = "TULSA"
	CreateSessionRequestUseProxyGeolocationCityTunis                   CreateSessionRequestUseProxyGeolocationCity = "TUNIS"
	CreateSessionRequestUseProxyGeolocationCityTunja                   CreateSessionRequestUseProxyGeolocationCity = "TUNJA"
	CreateSessionRequestUseProxyGeolocationCityTurin                   CreateSessionRequestUseProxyGeolocationCity = "TURIN"
	CreateSessionRequestUseProxyGeolocationCityTuxtlaGutierrez         CreateSessionRequestUseProxyGeolocationCity = "TUXTLA_GUTIERREZ"
	CreateSessionRequestUseProxyGeolocationCityTuzla                   CreateSessionRequestUseProxyGeolocationCity = "TUZLA"
	CreateSessionRequestUseProxyGeolocationCityUberaba                 CreateSessionRequestUseProxyGeolocationCity = "UBERABA"
	CreateSessionRequestUseProxyGeolocationCityUberlandia              CreateSessionRequestUseProxyGeolocationCity = "UBERLANDIA"
	CreateSessionRequestUseProxyGeolocationCityUfa                     CreateSessionRequestUseProxyGeolocationCity = "UFA"
	CreateSessionRequestUseProxyGeolocationCityUlanBator               CreateSessionRequestUseProxyGeolocationCity = "ULAN_BATOR"
	CreateSessionRequestUseProxyGeolocationCityUmeda                   CreateSessionRequestUseProxyGeolocationCity = "UMEDA"
	CreateSessionRequestUseProxyGeolocationCityUrdaneta                CreateSessionRequestUseProxyGeolocationCity = "URDANETA"
	CreateSessionRequestUseProxyGeolocationCityUsak                    CreateSessionRequestUseProxyGeolocationCity = "USAK"
	CreateSessionRequestUseProxyGeolocationCityVadodara                CreateSessionRequestUseProxyGeolocationCity = "VADODARA"
	CreateSessionRequestUseProxyGeolocationCityValencia                CreateSessionRequestUseProxyGeolocationCity = "VALENCIA"
	CreateSessionRequestUseProxyGeolocationCityValinhos                CreateSessionRequestUseProxyGeolocationCity = "VALINHOS"
	CreateSessionRequestUseProxyGeolocationCityValladolid              CreateSessionRequestUseProxyGeolocationCity = "VALLADOLID"
	CreateSessionRequestUseProxyGeolocationCityValledupar              CreateSessionRequestUseProxyGeolocationCity = "VALLEDUPAR"
	CreateSessionRequestUseProxyGeolocationCityValparaiso              CreateSessionRequestUseProxyGeolocationCity = "VALPARAISO"
	CreateSessionRequestUseProxyGeolocationCityValparaisoDeGoias       CreateSessionRequestUseProxyGeolocationCity = "VALPARAISO_DE GOIAS"
	CreateSessionRequestUseProxyGeolocationCityVan                     CreateSessionRequestUseProxyGeolocationCity = "VAN"
	CreateSessionRequestUseProxyGeolocationCityVancouver               CreateSessionRequestUseProxyGeolocationCity = "VANCOUVER"
	CreateSessionRequestUseProxyGeolocationCityVaranasi                CreateSessionRequestUseProxyGeolocationCity = "VARANASI"
	CreateSessionRequestUseProxyGeolocationCityVarginha                CreateSessionRequestUseProxyGeolocationCity = "VARGINHA"
	CreateSessionRequestUseProxyGeolocationCityVarna                   CreateSessionRequestUseProxyGeolocationCity = "VARNA"
	CreateSessionRequestUseProxyGeolocationCityVarzeaPaulista          CreateSessionRequestUseProxyGeolocationCity = "VARZEA_PAULISTA"
	CreateSessionRequestUseProxyGeolocationCityVenustianoCarranza      CreateSessionRequestUseProxyGeolocationCity = "VENUSTIANO_CARRANZA"
	CreateSessionRequestUseProxyGeolocationCityVeracruz                CreateSessionRequestUseProxyGeolocationCity = "VERACRUZ"
	CreateSessionRequestUseProxyGeolocationCityVerona                  CreateSessionRequestUseProxyGeolocationCity = "VERONA"
	CreateSessionRequestUseProxyGeolocationCityViamao                  CreateSessionRequestUseProxyGeolocationCity = "VIAMAO"
	CreateSessionRequestUseProxyGeolocationCityVictoria                CreateSessionRequestUseProxyGeolocationCity = "VICTORIA"
	CreateSessionRequestUseProxyGeolocationCityVienna                  CreateSessionRequestUseProxyGeolocationCity = "VIENNA"
	CreateSessionRequestUseProxyGeolocationCityVientiane               CreateSessionRequestUseProxyGeolocationCity = "VIENTIANE"
	CreateSessionRequestUseProxyGeolocationCityVigo                    CreateSessionRequestUseProxyGeolocationCity = "VIGO"
	CreateSessionRequestUseProxyGeolocationCityVijayawada              CreateSessionRequestUseProxyGeolocationCity = "VIJAYAWADA"
	CreateSessionRequestUseProxyGeolocationCityVilaNovaDeGaia          CreateSessionRequestUseProxyGeolocationCity = "VILA_NOVA DE GAIA"
	CreateSessionRequestUseProxyGeolocationCityVilaVelha               CreateSessionRequestUseProxyGeolocationCity = "VILA_VELHA"
	CreateSessionRequestUseProxyGeolocationCityVillaBallester          CreateSessionRequestUseProxyGeolocationCity = "VILLA_BALLESTER"
	CreateSessionRequestUseProxyGeolocationCityVillavicencio           CreateSessionRequestUseProxyGeolocationCity = "VILLAVICENCIO"
	CreateSessionRequestUseProxyGeolocationCityVilnius                 CreateSessionRequestUseProxyGeolocationCity = "VILNIUS"
	CreateSessionRequestUseProxyGeolocationCityVinaDelMar              CreateSessionRequestUseProxyGeolocationCity = "VINA_DEL MAR"
	CreateSessionRequestUseProxyGeolocationCityVinnytsia               CreateSessionRequestUseProxyGeolocationCity = "VINNYTSIA"
	CreateSessionRequestUseProxyGeolocationCityVirginiaBeach           CreateSessionRequestUseProxyGeolocationCity = "VIRGINIA_BEACH"
	CreateSessionRequestUseProxyGeolocationCityVisakhapatnam           CreateSessionRequestUseProxyGeolocationCity = "VISAKHAPATNAM"
	CreateSessionRequestUseProxyGeolocationCityVitoria                 CreateSessionRequestUseProxyGeolocationCity = "VITORIA"
	CreateSessionRequestUseProxyGeolocationCityVitoriaDaConquista      CreateSessionRequestUseProxyGeolocationCity = "VITORIA_DA CONQUISTA"
	CreateSessionRequestUseProxyGeolocationCityVitoriaDeSantoAntao     CreateSessionRequestUseProxyGeolocationCity = "VITORIA_DE SANTO ANTAO"
	CreateSessionRequestUseProxyGeolocationCityVoltaRedonda            CreateSessionRequestUseProxyGeolocationCity = "VOLTA_REDONDA"
	CreateSessionRequestUseProxyGeolocationCityVoronezh                CreateSessionRequestUseProxyGeolocationCity = "VORONEZH"
	CreateSessionRequestUseProxyGeolocationCityWarsaw                  CreateSessionRequestUseProxyGeolocationCity = "WARSAW"
	CreateSessionRequestUseProxyGeolocationCityWashington              CreateSessionRequestUseProxyGeolocationCity = "WASHINGTON"
	CreateSessionRequestUseProxyGeolocationCityWellington              CreateSessionRequestUseProxyGeolocationCity = "WELLINGTON"
	CreateSessionRequestUseProxyGeolocationCityWestPalmBeach           CreateSessionRequestUseProxyGeolocationCity = "WEST_PALM BEACH"
	CreateSessionRequestUseProxyGeolocationCityWichita                 CreateSessionRequestUseProxyGeolocationCity = "WICHITA"
	CreateSessionRequestUseProxyGeolocationCityWillemstad              CreateSessionRequestUseProxyGeolocationCity = "WILLEMSTAD"
	CreateSessionRequestUseProxyGeolocationCityWilmington              CreateSessionRequestUseProxyGeolocationCity = "WILMINGTON"
	CreateSessionRequestUseProxyGeolocationCityWindhoek                CreateSessionRequestUseProxyGeolocationCity = "WINDHOEK"
	CreateSessionRequestUseProxyGeolocationCityWindsor                 CreateSessionRequestUseProxyGeolocationCity = "WINDSOR"
	CreateSessionRequestUseProxyGeolocationCityWinnipeg                CreateSessionRequestUseProxyGeolocationCity = "WINNIPEG"
	CreateSessionRequestUseProxyGeolocationCityWolverhampton           CreateSessionRequestUseProxyGeolocationCity = "WOLVERHAMPTON"
	CreateSessionRequestUseProxyGeolocationCityWoodbridge              CreateSessionRequestUseProxyGeolocationCity = "WOODBRIDGE"
	CreateSessionRequestUseProxyGeolocationCityWroclaw                 CreateSessionRequestUseProxyGeolocationCity = "WROCLAW"
	CreateSessionRequestUseProxyGeolocationCityWuppertal               CreateSessionRequestUseProxyGeolocationCity = "WUPPERTAL"
	CreateSessionRequestUseProxyGeolocationCityXalapa                  CreateSessionRequestUseProxyGeolocationCity = "XALAPA"
	CreateSessionRequestUseProxyGeolocationCityYalova                  CreateSessionRequestUseProxyGeolocationCity = "YALOVA"
	CreateSessionRequestUseProxyGeolocationCityYangon                  CreateSessionRequestUseProxyGeolocationCity = "YANGON"
	CreateSessionRequestUseProxyGeolocationCityYekaterinburg           CreateSessionRequestUseProxyGeolocationCity = "YEKATERINBURG"
	CreateSessionRequestUseProxyGeolocationCityYerevan                 CreateSessionRequestUseProxyGeolocationCity = "YEREVAN"
	CreateSessionRequestUseProxyGeolocationCityYogyakarta              CreateSessionRequestUseProxyGeolocationCity = "YOGYAKARTA"
	CreateSessionRequestUseProxyGeolocationCityYokohama                CreateSessionRequestUseProxyGeolocationCity = "YOKOHAMA"
	CreateSessionRequestUseProxyGeolocationCityYonginSi                CreateSessionRequestUseProxyGeolocationCity = "YONGIN_SI"
	CreateSessionRequestUseProxyGeolocationCityZabrze                  CreateSessionRequestUseProxyGeolocationCity = "ZABRZE"
	CreateSessionRequestUseProxyGeolocationCityZagazig                 CreateSessionRequestUseProxyGeolocationCity = "ZAGAZIG"
	CreateSessionRequestUseProxyGeolocationCityZagreb                  CreateSessionRequestUseProxyGeolocationCity = "ZAGREB"
	CreateSessionRequestUseProxyGeolocationCityZamboangaCity           CreateSessionRequestUseProxyGeolocationCity = "ZAMBOANGA_CITY"
	CreateSessionRequestUseProxyGeolocationCityZapopan                 CreateSessionRequestUseProxyGeolocationCity = "ZAPOPAN"
	CreateSessionRequestUseProxyGeolocationCityZaporizhzhya            CreateSessionRequestUseProxyGeolocationCity = "ZAPORIZHZHYA"
	CreateSessionRequestUseProxyGeolocationCityZaragoza                CreateSessionRequestUseProxyGeolocationCity = "ZARAGOZA"
	CreateSessionRequestUseProxyGeolocationCityZhongliDistrict         CreateSessionRequestUseProxyGeolocationCity = "ZHONGLI_DISTRICT"
	CreateSessionRequestUseProxyGeolocationCityZielonaGora             CreateSessionRequestUseProxyGeolocationCity = "ZIELONA_GORA"
	CreateSessionRequestUseProxyGeolocationCityZonguldak               CreateSessionRequestUseProxyGeolocationCity = "ZONGULDAK"
	CreateSessionRequestUseProxyGeolocationCityZurich                  CreateSessionRequestUseProxyGeolocationCity = "ZURICH"
)

type CreateSessionRequestUseProxyGeolocationCountry string

const (
	CreateSessionRequestUseProxyGeolocationCountryUs CreateSessionRequestUseProxyGeolocationCountry = "US"
	CreateSessionRequestUseProxyGeolocationCountryCa CreateSessionRequestUseProxyGeolocationCountry = "CA"
	CreateSessionRequestUseProxyGeolocationCountryMx CreateSessionRequestUseProxyGeolocationCountry = "MX"
	CreateSessionRequestUseProxyGeolocationCountryGB CreateSessionRequestUseProxyGeolocationCountry = "GB"
	CreateSessionRequestUseProxyGeolocationCountryDe CreateSessionRequestUseProxyGeolocationCountry = "DE"
	CreateSessionRequestUseProxyGeolocationCountryFr CreateSessionRequestUseProxyGeolocationCountry = "FR"
	CreateSessionRequestUseProxyGeolocationCountryIt CreateSessionRequestUseProxyGeolocationCountry = "IT"
	CreateSessionRequestUseProxyGeolocationCountryEs CreateSessionRequestUseProxyGeolocationCountry = "ES"
	CreateSessionRequestUseProxyGeolocationCountryPl CreateSessionRequestUseProxyGeolocationCountry = "PL"
	CreateSessionRequestUseProxyGeolocationCountryNl CreateSessionRequestUseProxyGeolocationCountry = "NL"
	CreateSessionRequestUseProxyGeolocationCountrySe CreateSessionRequestUseProxyGeolocationCountry = "SE"
	CreateSessionRequestUseProxyGeolocationCountryNo CreateSessionRequestUseProxyGeolocationCountry = "NO"
	CreateSessionRequestUseProxyGeolocationCountryDk CreateSessionRequestUseProxyGeolocationCountry = "DK"
	CreateSessionRequestUseProxyGeolocationCountryFi CreateSessionRequestUseProxyGeolocationCountry = "FI"
	CreateSessionRequestUseProxyGeolocationCountryCh CreateSessionRequestUseProxyGeolocationCountry = "CH"
	CreateSessionRequestUseProxyGeolocationCountryAt CreateSessionRequestUseProxyGeolocationCountry = "AT"
	CreateSessionRequestUseProxyGeolocationCountryBe CreateSessionRequestUseProxyGeolocationCountry = "BE"
	CreateSessionRequestUseProxyGeolocationCountryIe CreateSessionRequestUseProxyGeolocationCountry = "IE"
	CreateSessionRequestUseProxyGeolocationCountryPt CreateSessionRequestUseProxyGeolocationCountry = "PT"
	CreateSessionRequestUseProxyGeolocationCountryGr CreateSessionRequestUseProxyGeolocationCountry = "GR"
	CreateSessionRequestUseProxyGeolocationCountryCz CreateSessionRequestUseProxyGeolocationCountry = "CZ"
	CreateSessionRequestUseProxyGeolocationCountryHu CreateSessionRequestUseProxyGeolocationCountry = "HU"
	CreateSessionRequestUseProxyGeolocationCountryRo CreateSessionRequestUseProxyGeolocationCountry = "RO"
	CreateSessionRequestUseProxyGeolocationCountryBg CreateSessionRequestUseProxyGeolocationCountry = "BG"
	CreateSessionRequestUseProxyGeolocationCountrySk CreateSessionRequestUseProxyGeolocationCountry = "SK"
	CreateSessionRequestUseProxyGeolocationCountrySi CreateSessionRequestUseProxyGeolocationCountry = "SI"
	CreateSessionRequestUseProxyGeolocationCountryHr CreateSessionRequestUseProxyGeolocationCountry = "HR"
	CreateSessionRequestUseProxyGeolocationCountryEe CreateSessionRequestUseProxyGeolocationCountry = "EE"
	CreateSessionRequestUseProxyGeolocationCountryLv CreateSessionRequestUseProxyGeolocationCountry = "LV"
	CreateSessionRequestUseProxyGeolocationCountryLt CreateSessionRequestUseProxyGeolocationCountry = "LT"
	CreateSessionRequestUseProxyGeolocationCountryLu CreateSessionRequestUseProxyGeolocationCountry = "LU"
	CreateSessionRequestUseProxyGeolocationCountryMt CreateSessionRequestUseProxyGeolocationCountry = "MT"
	CreateSessionRequestUseProxyGeolocationCountryCy CreateSessionRequestUseProxyGeolocationCountry = "CY"
	CreateSessionRequestUseProxyGeolocationCountryIs CreateSessionRequestUseProxyGeolocationCountry = "IS"
	CreateSessionRequestUseProxyGeolocationCountryLi CreateSessionRequestUseProxyGeolocationCountry = "LI"
	CreateSessionRequestUseProxyGeolocationCountryMc CreateSessionRequestUseProxyGeolocationCountry = "MC"
	CreateSessionRequestUseProxyGeolocationCountrySm CreateSessionRequestUseProxyGeolocationCountry = "SM"
	CreateSessionRequestUseProxyGeolocationCountryVa CreateSessionRequestUseProxyGeolocationCountry = "VA"
	CreateSessionRequestUseProxyGeolocationCountryJp CreateSessionRequestUseProxyGeolocationCountry = "JP"
	CreateSessionRequestUseProxyGeolocationCountryKr CreateSessionRequestUseProxyGeolocationCountry = "KR"
	CreateSessionRequestUseProxyGeolocationCountryCn CreateSessionRequestUseProxyGeolocationCountry = "CN"
	CreateSessionRequestUseProxyGeolocationCountryHk CreateSessionRequestUseProxyGeolocationCountry = "HK"
	CreateSessionRequestUseProxyGeolocationCountryTw CreateSessionRequestUseProxyGeolocationCountry = "TW"
	CreateSessionRequestUseProxyGeolocationCountrySg CreateSessionRequestUseProxyGeolocationCountry = "SG"
	CreateSessionRequestUseProxyGeolocationCountryAu CreateSessionRequestUseProxyGeolocationCountry = "AU"
	CreateSessionRequestUseProxyGeolocationCountryNz CreateSessionRequestUseProxyGeolocationCountry = "NZ"
	CreateSessionRequestUseProxyGeolocationCountryIn CreateSessionRequestUseProxyGeolocationCountry = "IN"
	CreateSessionRequestUseProxyGeolocationCountryTh CreateSessionRequestUseProxyGeolocationCountry = "TH"
	CreateSessionRequestUseProxyGeolocationCountryMy CreateSessionRequestUseProxyGeolocationCountry = "MY"
	CreateSessionRequestUseProxyGeolocationCountryPh CreateSessionRequestUseProxyGeolocationCountry = "PH"
	CreateSessionRequestUseProxyGeolocationCountryID CreateSessionRequestUseProxyGeolocationCountry = "ID"
	CreateSessionRequestUseProxyGeolocationCountryVn CreateSessionRequestUseProxyGeolocationCountry = "VN"
	CreateSessionRequestUseProxyGeolocationCountryAf CreateSessionRequestUseProxyGeolocationCountry = "AF"
	CreateSessionRequestUseProxyGeolocationCountryBd CreateSessionRequestUseProxyGeolocationCountry = "BD"
	CreateSessionRequestUseProxyGeolocationCountryBn CreateSessionRequestUseProxyGeolocationCountry = "BN"
	CreateSessionRequestUseProxyGeolocationCountryKh CreateSessionRequestUseProxyGeolocationCountry = "KH"
	CreateSessionRequestUseProxyGeolocationCountryLa CreateSessionRequestUseProxyGeolocationCountry = "LA"
	CreateSessionRequestUseProxyGeolocationCountryLk CreateSessionRequestUseProxyGeolocationCountry = "LK"
	CreateSessionRequestUseProxyGeolocationCountryMm CreateSessionRequestUseProxyGeolocationCountry = "MM"
	CreateSessionRequestUseProxyGeolocationCountryNp CreateSessionRequestUseProxyGeolocationCountry = "NP"
	CreateSessionRequestUseProxyGeolocationCountryPk CreateSessionRequestUseProxyGeolocationCountry = "PK"
	CreateSessionRequestUseProxyGeolocationCountryFj CreateSessionRequestUseProxyGeolocationCountry = "FJ"
	CreateSessionRequestUseProxyGeolocationCountryPg CreateSessionRequestUseProxyGeolocationCountry = "PG"
	CreateSessionRequestUseProxyGeolocationCountryAe CreateSessionRequestUseProxyGeolocationCountry = "AE"
	CreateSessionRequestUseProxyGeolocationCountrySa CreateSessionRequestUseProxyGeolocationCountry = "SA"
	CreateSessionRequestUseProxyGeolocationCountryIl CreateSessionRequestUseProxyGeolocationCountry = "IL"
	CreateSessionRequestUseProxyGeolocationCountryTr CreateSessionRequestUseProxyGeolocationCountry = "TR"
	CreateSessionRequestUseProxyGeolocationCountryIr CreateSessionRequestUseProxyGeolocationCountry = "IR"
	CreateSessionRequestUseProxyGeolocationCountryIq CreateSessionRequestUseProxyGeolocationCountry = "IQ"
	CreateSessionRequestUseProxyGeolocationCountryJo CreateSessionRequestUseProxyGeolocationCountry = "JO"
	CreateSessionRequestUseProxyGeolocationCountryKw CreateSessionRequestUseProxyGeolocationCountry = "KW"
	CreateSessionRequestUseProxyGeolocationCountryLb CreateSessionRequestUseProxyGeolocationCountry = "LB"
	CreateSessionRequestUseProxyGeolocationCountryOm CreateSessionRequestUseProxyGeolocationCountry = "OM"
	CreateSessionRequestUseProxyGeolocationCountryQa CreateSessionRequestUseProxyGeolocationCountry = "QA"
	CreateSessionRequestUseProxyGeolocationCountryBh CreateSessionRequestUseProxyGeolocationCountry = "BH"
	CreateSessionRequestUseProxyGeolocationCountryYe CreateSessionRequestUseProxyGeolocationCountry = "YE"
	CreateSessionRequestUseProxyGeolocationCountrySy CreateSessionRequestUseProxyGeolocationCountry = "SY"
	CreateSessionRequestUseProxyGeolocationCountryZa CreateSessionRequestUseProxyGeolocationCountry = "ZA"
	CreateSessionRequestUseProxyGeolocationCountryEg CreateSessionRequestUseProxyGeolocationCountry = "EG"
	CreateSessionRequestUseProxyGeolocationCountryMa CreateSessionRequestUseProxyGeolocationCountry = "MA"
	CreateSessionRequestUseProxyGeolocationCountryNg CreateSessionRequestUseProxyGeolocationCountry = "NG"
	CreateSessionRequestUseProxyGeolocationCountryKe CreateSessionRequestUseProxyGeolocationCountry = "KE"
	CreateSessionRequestUseProxyGeolocationCountryDz CreateSessionRequestUseProxyGeolocationCountry = "DZ"
	CreateSessionRequestUseProxyGeolocationCountryAo CreateSessionRequestUseProxyGeolocationCountry = "AO"
	CreateSessionRequestUseProxyGeolocationCountryBw CreateSessionRequestUseProxyGeolocationCountry = "BW"
	CreateSessionRequestUseProxyGeolocationCountryEt CreateSessionRequestUseProxyGeolocationCountry = "ET"
	CreateSessionRequestUseProxyGeolocationCountryGh CreateSessionRequestUseProxyGeolocationCountry = "GH"
	CreateSessionRequestUseProxyGeolocationCountryCi CreateSessionRequestUseProxyGeolocationCountry = "CI"
	CreateSessionRequestUseProxyGeolocationCountryLy CreateSessionRequestUseProxyGeolocationCountry = "LY"
	CreateSessionRequestUseProxyGeolocationCountryMz CreateSessionRequestUseProxyGeolocationCountry = "MZ"
	CreateSessionRequestUseProxyGeolocationCountryRw CreateSessionRequestUseProxyGeolocationCountry = "RW"
	CreateSessionRequestUseProxyGeolocationCountrySn CreateSessionRequestUseProxyGeolocationCountry = "SN"
	CreateSessionRequestUseProxyGeolocationCountryTn CreateSessionRequestUseProxyGeolocationCountry = "TN"
	CreateSessionRequestUseProxyGeolocationCountryUg CreateSessionRequestUseProxyGeolocationCountry = "UG"
	CreateSessionRequestUseProxyGeolocationCountryZm CreateSessionRequestUseProxyGeolocationCountry = "ZM"
	CreateSessionRequestUseProxyGeolocationCountryZw CreateSessionRequestUseProxyGeolocationCountry = "ZW"
	CreateSessionRequestUseProxyGeolocationCountryTz CreateSessionRequestUseProxyGeolocationCountry = "TZ"
	CreateSessionRequestUseProxyGeolocationCountryMu CreateSessionRequestUseProxyGeolocationCountry = "MU"
	CreateSessionRequestUseProxyGeolocationCountrySc CreateSessionRequestUseProxyGeolocationCountry = "SC"
	CreateSessionRequestUseProxyGeolocationCountryBr CreateSessionRequestUseProxyGeolocationCountry = "BR"
	CreateSessionRequestUseProxyGeolocationCountryAr CreateSessionRequestUseProxyGeolocationCountry = "AR"
	CreateSessionRequestUseProxyGeolocationCountryCl CreateSessionRequestUseProxyGeolocationCountry = "CL"
	CreateSessionRequestUseProxyGeolocationCountryCo CreateSessionRequestUseProxyGeolocationCountry = "CO"
	CreateSessionRequestUseProxyGeolocationCountryPe CreateSessionRequestUseProxyGeolocationCountry = "PE"
	CreateSessionRequestUseProxyGeolocationCountryVe CreateSessionRequestUseProxyGeolocationCountry = "VE"
	CreateSessionRequestUseProxyGeolocationCountryEc CreateSessionRequestUseProxyGeolocationCountry = "EC"
	CreateSessionRequestUseProxyGeolocationCountryUy CreateSessionRequestUseProxyGeolocationCountry = "UY"
	CreateSessionRequestUseProxyGeolocationCountryPy CreateSessionRequestUseProxyGeolocationCountry = "PY"
	CreateSessionRequestUseProxyGeolocationCountryBo CreateSessionRequestUseProxyGeolocationCountry = "BO"
	CreateSessionRequestUseProxyGeolocationCountryCr CreateSessionRequestUseProxyGeolocationCountry = "CR"
	CreateSessionRequestUseProxyGeolocationCountryCu CreateSessionRequestUseProxyGeolocationCountry = "CU"
	CreateSessionRequestUseProxyGeolocationCountryDo CreateSessionRequestUseProxyGeolocationCountry = "DO"
	CreateSessionRequestUseProxyGeolocationCountryGt CreateSessionRequestUseProxyGeolocationCountry = "GT"
	CreateSessionRequestUseProxyGeolocationCountryHn CreateSessionRequestUseProxyGeolocationCountry = "HN"
	CreateSessionRequestUseProxyGeolocationCountryJm CreateSessionRequestUseProxyGeolocationCountry = "JM"
	CreateSessionRequestUseProxyGeolocationCountryNi CreateSessionRequestUseProxyGeolocationCountry = "NI"
	CreateSessionRequestUseProxyGeolocationCountryPa CreateSessionRequestUseProxyGeolocationCountry = "PA"
	CreateSessionRequestUseProxyGeolocationCountrySv CreateSessionRequestUseProxyGeolocationCountry = "SV"
	CreateSessionRequestUseProxyGeolocationCountryTt CreateSessionRequestUseProxyGeolocationCountry = "TT"
	CreateSessionRequestUseProxyGeolocationCountryBb CreateSessionRequestUseProxyGeolocationCountry = "BB"
	CreateSessionRequestUseProxyGeolocationCountryBz CreateSessionRequestUseProxyGeolocationCountry = "BZ"
	CreateSessionRequestUseProxyGeolocationCountryGy CreateSessionRequestUseProxyGeolocationCountry = "GY"
	CreateSessionRequestUseProxyGeolocationCountrySr CreateSessionRequestUseProxyGeolocationCountry = "SR"
	CreateSessionRequestUseProxyGeolocationCountryRu CreateSessionRequestUseProxyGeolocationCountry = "RU"
	CreateSessionRequestUseProxyGeolocationCountryUa CreateSessionRequestUseProxyGeolocationCountry = "UA"
	CreateSessionRequestUseProxyGeolocationCountryBy CreateSessionRequestUseProxyGeolocationCountry = "BY"
	CreateSessionRequestUseProxyGeolocationCountryKz CreateSessionRequestUseProxyGeolocationCountry = "KZ"
	CreateSessionRequestUseProxyGeolocationCountryUz CreateSessionRequestUseProxyGeolocationCountry = "UZ"
	CreateSessionRequestUseProxyGeolocationCountryAz CreateSessionRequestUseProxyGeolocationCountry = "AZ"
	CreateSessionRequestUseProxyGeolocationCountryGe CreateSessionRequestUseProxyGeolocationCountry = "GE"
	CreateSessionRequestUseProxyGeolocationCountryAm CreateSessionRequestUseProxyGeolocationCountry = "AM"
	CreateSessionRequestUseProxyGeolocationCountryMd CreateSessionRequestUseProxyGeolocationCountry = "MD"
	CreateSessionRequestUseProxyGeolocationCountryMk CreateSessionRequestUseProxyGeolocationCountry = "MK"
	CreateSessionRequestUseProxyGeolocationCountryAl CreateSessionRequestUseProxyGeolocationCountry = "AL"
	CreateSessionRequestUseProxyGeolocationCountryBa CreateSessionRequestUseProxyGeolocationCountry = "BA"
	CreateSessionRequestUseProxyGeolocationCountryRs CreateSessionRequestUseProxyGeolocationCountry = "RS"
	CreateSessionRequestUseProxyGeolocationCountryMe CreateSessionRequestUseProxyGeolocationCountry = "ME"
	CreateSessionRequestUseProxyGeolocationCountryXk CreateSessionRequestUseProxyGeolocationCountry = "XK"
	CreateSessionRequestUseProxyGeolocationCountryMn CreateSessionRequestUseProxyGeolocationCountry = "MN"
	CreateSessionRequestUseProxyGeolocationCountryKg CreateSessionRequestUseProxyGeolocationCountry = "KG"
	CreateSessionRequestUseProxyGeolocationCountryTj CreateSessionRequestUseProxyGeolocationCountry = "TJ"
	CreateSessionRequestUseProxyGeolocationCountryTm CreateSessionRequestUseProxyGeolocationCountry = "TM"
)

type CreateSessionRequestUseProxyGeolocationState string

const (
	CreateSessionRequestUseProxyGeolocationStateAl CreateSessionRequestUseProxyGeolocationState = "AL"
	CreateSessionRequestUseProxyGeolocationStateAk CreateSessionRequestUseProxyGeolocationState = "AK"
	CreateSessionRequestUseProxyGeolocationStateAz CreateSessionRequestUseProxyGeolocationState = "AZ"
	CreateSessionRequestUseProxyGeolocationStateAr CreateSessionRequestUseProxyGeolocationState = "AR"
	CreateSessionRequestUseProxyGeolocationStateCa CreateSessionRequestUseProxyGeolocationState = "CA"
	CreateSessionRequestUseProxyGeolocationStateCo CreateSessionRequestUseProxyGeolocationState = "CO"
	CreateSessionRequestUseProxyGeolocationStateCt CreateSessionRequestUseProxyGeolocationState = "CT"
	CreateSessionRequestUseProxyGeolocationStateDe CreateSessionRequestUseProxyGeolocationState = "DE"
	CreateSessionRequestUseProxyGeolocationStateFl CreateSessionRequestUseProxyGeolocationState = "FL"
	CreateSessionRequestUseProxyGeolocationStateGa CreateSessionRequestUseProxyGeolocationState = "GA"
	CreateSessionRequestUseProxyGeolocationStateHi CreateSessionRequestUseProxyGeolocationState = "HI"
	CreateSessionRequestUseProxyGeolocationStateID CreateSessionRequestUseProxyGeolocationState = "ID"
	CreateSessionRequestUseProxyGeolocationStateIl CreateSessionRequestUseProxyGeolocationState = "IL"
	CreateSessionRequestUseProxyGeolocationStateIn CreateSessionRequestUseProxyGeolocationState = "IN"
	CreateSessionRequestUseProxyGeolocationStateIa CreateSessionRequestUseProxyGeolocationState = "IA"
	CreateSessionRequestUseProxyGeolocationStateKs CreateSessionRequestUseProxyGeolocationState = "KS"
	CreateSessionRequestUseProxyGeolocationStateKy CreateSessionRequestUseProxyGeolocationState = "KY"
	CreateSessionRequestUseProxyGeolocationStateLa CreateSessionRequestUseProxyGeolocationState = "LA"
	CreateSessionRequestUseProxyGeolocationStateMe CreateSessionRequestUseProxyGeolocationState = "ME"
	CreateSessionRequestUseProxyGeolocationStateMd CreateSessionRequestUseProxyGeolocationState = "MD"
	CreateSessionRequestUseProxyGeolocationStateMa CreateSessionRequestUseProxyGeolocationState = "MA"
	CreateSessionRequestUseProxyGeolocationStateMi CreateSessionRequestUseProxyGeolocationState = "MI"
	CreateSessionRequestUseProxyGeolocationStateMn CreateSessionRequestUseProxyGeolocationState = "MN"
	CreateSessionRequestUseProxyGeolocationStateMs CreateSessionRequestUseProxyGeolocationState = "MS"
	CreateSessionRequestUseProxyGeolocationStateMo CreateSessionRequestUseProxyGeolocationState = "MO"
	CreateSessionRequestUseProxyGeolocationStateMt CreateSessionRequestUseProxyGeolocationState = "MT"
	CreateSessionRequestUseProxyGeolocationStateNe CreateSessionRequestUseProxyGeolocationState = "NE"
	CreateSessionRequestUseProxyGeolocationStateNv CreateSessionRequestUseProxyGeolocationState = "NV"
	CreateSessionRequestUseProxyGeolocationStateNh CreateSessionRequestUseProxyGeolocationState = "NH"
	CreateSessionRequestUseProxyGeolocationStateNj CreateSessionRequestUseProxyGeolocationState = "NJ"
	CreateSessionRequestUseProxyGeolocationStateNm CreateSessionRequestUseProxyGeolocationState = "NM"
	CreateSessionRequestUseProxyGeolocationStateNy CreateSessionRequestUseProxyGeolocationState = "NY"
	CreateSessionRequestUseProxyGeolocationStateNc CreateSessionRequestUseProxyGeolocationState = "NC"
	CreateSessionRequestUseProxyGeolocationStateNd CreateSessionRequestUseProxyGeolocationState = "ND"
	CreateSessionRequestUseProxyGeolocationStateOh CreateSessionRequestUseProxyGeolocationState = "OH"
	CreateSessionRequestUseProxyGeolocationStateOk CreateSessionRequestUseProxyGeolocationState = "OK"
	CreateSessionRequestUseProxyGeolocationStateOr CreateSessionRequestUseProxyGeolocationState = "OR"
	CreateSessionRequestUseProxyGeolocationStatePa CreateSessionRequestUseProxyGeolocationState = "PA"
	CreateSessionRequestUseProxyGeolocationStateRi CreateSessionRequestUseProxyGeolocationState = "RI"
	CreateSessionRequestUseProxyGeolocationStateSc CreateSessionRequestUseProxyGeolocationState = "SC"
	CreateSessionRequestUseProxyGeolocationStateSd CreateSessionRequestUseProxyGeolocationState = "SD"
	CreateSessionRequestUseProxyGeolocationStateTn CreateSessionRequestUseProxyGeolocationState = "TN"
	CreateSessionRequestUseProxyGeolocationStateTx CreateSessionRequestUseProxyGeolocationState = "TX"
	CreateSessionRequestUseProxyGeolocationStateUt CreateSessionRequestUseProxyGeolocationState = "UT"
	CreateSessionRequestUseProxyGeolocationStateVt CreateSessionRequestUseProxyGeolocationState = "VT"
	CreateSessionRequestUseProxyGeolocationStateVa CreateSessionRequestUseProxyGeolocationState = "VA"
	CreateSessionRequestUseProxyGeolocationStateWa CreateSessionRequestUseProxyGeolocationState = "WA"
	CreateSessionRequestUseProxyGeolocationStateWv CreateSessionRequestUseProxyGeolocationState = "WV"
	CreateSessionRequestUseProxyGeolocationStateWi CreateSessionRequestUseProxyGeolocationState = "WI"
	CreateSessionRequestUseProxyGeolocationStateWy CreateSessionRequestUseProxyGeolocationState = "WY"
	CreateSessionRequestUseProxyGeolocationStateDc CreateSessionRequestUseProxyGeolocationState = "DC"
	CreateSessionRequestUseProxyGeolocationStatePr CreateSessionRequestUseProxyGeolocationState = "PR"
	CreateSessionRequestUseProxyGeolocationStateGu CreateSessionRequestUseProxyGeolocationState = "GU"
	CreateSessionRequestUseProxyGeolocationStateVi CreateSessionRequestUseProxyGeolocationState = "VI"
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

type FileList struct {
	// Array of files for the current page
	Data []FileListData `json:"data" api:"required"`
	JSON fileListJSON   `json:"-"`
}

// fileListJSON contains the JSON metadata for the struct [FileList]
type fileListJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileListJSON) RawJSON() string {
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

// FileListData Array of files for the current page
type FileListData struct {
	// Timestamp when the file was created
	LastModified time.Time `json:"lastModified" api:"required"`
	// Path to the file in the storage system
	Path string `json:"path" api:"required"`
	// Size of the file in bytes
	Size int64            `json:"size" api:"required"`
	JSON fileListDataJSON `json:"-"`
}

// fileListDataJSON contains the JSON metadata for the struct [FileListData]
type fileListDataJSON struct {
	LastModified apijson.Field
	Path         apijson.Field
	Size         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FileListData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileListDataJSON) RawJSON() string {
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
