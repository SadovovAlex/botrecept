language-name = English (US)
language-flag = 🇺🇸
language-menu =
    <b>Current language:</b> { $languageFlag} { $languageName }

    <b>Select below the language you want to use the bot.</b>
language-changed = The language has been changed successfully.
measurement-unit = m
start-button = Start a chat.
start-message =
    Привет <b>{ $userFirstName }</b> — Я <b>{ $botName }</b>, и я здесь, чтобы помочь вам с рецептами и приготовлением блюд.
    Что я могу сделать для вас:
    <b>Поиск рецептов:</b> Найдите идеальные рецепты для любого случая — от простых закусок до сложных блюд.
    <b>Советы по приготовлению:</b> Получите полезные советы и рекомендации по техникам приготовления и выбору ингредиентов.
    <b>Магазин специй:</b> Ознакомьтесь с нашим ассортиментом специй, чтобы добавить уникальные ароматы и вкусы в ваши блюда. 
    Вы можете легко заказать необходимые специи прямо у нас! <a href='zzzzz'>Магазин</a>
start-message-group =
    Привет, Я <b>{ $botName }</b>
    У меня есть много функций для использования напиши мне в личный чат.
language-button = Язык
help-button = ❔Мои функции
about-button =  ℹ️ О нас
recept-button =  🍽 Рецепт
donation-button = 💵 Поддержка
news-channel-button = 📢 Канал
about-your-data-button = О персональных даннных
back-button = ↩️ Назад
denied-button-alert = This button is not for you.
privacy-policy-button = 🔒 Соглашение
privacy-policy-group = To acess privacy policy, <b>click on the button below.</b>

recept = Рецепты
recept-help = 
    <b>Запрос рецептов</b>
    Можно попросить бота сгенерировать рецепт, случайный или на основе ингредиентов, для этого необходимо дать команду боту
    <b>/recept картошка, грибы</b>
about =
    <b>— яяяяяя</b>
    His fame arose from an image in which he is seen next to a woman shouting angrily

    <b>- Source code:</b> <a href='zzzzzz'>zzzzz</a>
    <b>- Developer:</b> @wrwfx
    <i>This project.............</i>

    <b>💸 Donate: Help keep the bot online.</b>
    • Pix Key and PayPal Email: <code>zzzzzz</code>

    If you want to contribute in another way, with a card for example, click on the button below and you will be redirected to the Ko-Fi donation link.
privacy-policy-private =
    <b>Privacy Policy.</b>

    is built to provide users with transparency and trust. 
    Thank you for your trust, and I am fully dedicated to protecting your privacy.

    This privacy policy may be updated, and any changes will be communicated through Channel - @zzzzzz.
about-your-data = 
    <b>About your data.</b>

    <b>1. Data collection.</b>
    The bot only collects essential information to provide a personalized experience.
    <b>The data collected includes:</b>
    - <b>Your Telegram user information:</b> User ID, first name, language, and username.
    - <b>Your settings :</b> Settings you have configured in the bot, such as your language and LastFM username,  all as provided by the user themselves.

    <b>2. Data usage.</b>
    The data collected by the bot is used exclusively to enhance the user experience and provide a more efficient service.
    - <b>Your Telegram user information</b> is used for identification and communication with the user.
    - <b>Your settings </b> are used to integrate and personalize the bot's services.

    <b>3. Data Sharing.</b>
    The data collected by the bot is not shared with third parties, except where required by law. 
    All your data is stored securely.

    <b>Note:</b> Your Telegram user information is public information from the platform. We do not know your "real" data.
help =
    Это мои функциональные модули
    <b>Нажмите на кнопку для ознакомления с каждым.</b>

    <b>Примечание:</b>
    Некоторые модули поддерживают конфигурацию для администраторов групп.
relative-duration-seconds = { $count ->
    [one] { $count } second
    *[other] { $count } seconds
}
relative-duration-minutes = { $count ->
    [one] { $count } minute
    *[other] { $count } minutes
}
relative-duration-hours = { $count ->
    [one] { $count } hour
    *[other] { $count } hours
}
relative-duration-days = { $count ->
    [one] { $count } day
    *[other] { $count } days
}
relative-duration-weeks = { $count ->
    [one] { $count } week
    *[other] { $count } weeks
}
relative-duration-months = { $count ->
    [one] { $count } month
    *[other] { $count } months
}
afk = AFK
afk-help = 
    <b>AFK — <i>Away from Keyboard</i></b>

    <b>AFK stands for</b> <i>“away from keyboard”</i>. It is basically Internet slang to say that you are away.

    <b>— Commands:</b>
    <b>/afk (reason):</b> Define yourself as away.
    <b>brb (reason):</b> Same as the afk command, but not a command; no need to use <code>/</code>.
user-now-unavailable = <b>{ $userFirstName }</b> is now unavailable!
user-unavailable =
    <b><a href='tg://user?id={ $userID }'>{ $userFirstName }</a></b> is <b>unavailable!</b>
    Last seen <code>{ $duration }</code> ago
user-unavailable-reason = <b>Reason:</b> <code>{ $reason }</code>
now-available = <b><a href='tg://user?id={ $userID }'>{ $userFirstName }</a></b> is back after <code>{ $duration }</code> away!
config = Settings
config-help =
    <b>Settings:</b>

    This module is designed to be <b>used in groups.</b>
    You must be an administrator to use them.

    <b>— Commands:</b>
    <b>/disable (command):</b> Disables the specified command in the group.
    <b>/enable (command):</b> Reactivates a command that was previously disabled.
    <b>/disableable:</b> Lists all commands that can be disabled.
    <b>/disabled:</b> Shows all commands that are currently disabled.
    <b>/config:</b> Opens a menu with group configuration options.
config-message =
    <b>Settings —</b> Here are my settings for this group.
    To know more, <b>click on the buttons below.</b>
config-medias =
    <b>Medias module settings:</b>
    To know more about the <b><i>medias</i></b> module, use /help in my private chat.

    <b>To know more about each configuration, click on its name.</b>
    <i>These settings are for this group only, not for other groups or private chats.</i>
caption-button = Captions:
automatic-button = Auto:
command-enabled = The command <code>{ $command }</code> has been successfully enabled.
command-already-enabled = The command <code>{ $command }</code> is already enabled.
enable-commands-usage =
    Please specify the command you want to enable. To see which commands are currently disabled, use /disabled.

    <b>Usage:</b> <code>/enable (command)</code>
no-disabled-commands = There are no disabled commands <b>in this group.</b>
disabled-commands = <b>Disabled commands:</b>
disableables-commands = <b>Disableable commands:</b>
command-already-disabled = The command <code>{ $command }</code> is already disabled.
command-disabled = The command <code>{ $command }</code> has been successfully disabled.
disable-commands-usage =
    Please specify the command you want to disable. To view the list of disableable commands, use /disableable.

    <b>Usage:</b> <code>/disable (command)</code>
command-not-deactivatable = The <code>{ $command }</code> command <b>cannot be deactivated.</b>
medias = Medias
medias-help =
    <b>Media Downloader</b>

    When sharing links on Telegram, some sites don't display an image or video preview. This module enables Smudge to automatically detect links from supported sites and send the videos and images contained within them.

    <b>Currently supported sites:</b> <i>Instagram</i>, <i>TikTok</i>, <i>Twitter/X</i>, <i>Threads</i>, <i>Reddit</i>, <i>Bluesky</i> <i>YouTube Shorts</i>, 

    <b>Note:</b> 
    This module contains additional settings for groups. 
    You can disable automatic downloads and captions in the groups.

    <b>— Commands:</b>
    <b>/dl | /sdl (link):</b> This command is for when you disable automatic downloads in groups.
    <b>/ytdl (link)</b>: Downloads videos from <b>YouTube</b>
    The maximum quality of video downloads is 1080p. You can also download just the audio of the video with this command.
youtube-no-url =
    You need to specify a valid YouTube link to download.
    <b>Example:</b> <code>/ytdl https://youtu.be/dQw4w9WgXcQ</code>
youtube-invalid-url = That YouTube link doesn't exist or it's a private video.
youtube-video-info =
    <b>Title:</b> { $title }
    <b>Uploader:</b> { $author }
    <b>Size:</b> <code>{ $audioSize }</code> (Audio) | <code>{ $videoSize }</code> (Video)
    <b>Duration:</b> { $duration }
youtube-download-video-button = Download video
youtube-download-audio-button = Download audio
video-exceeds-limit = 
    This video exceeds the size of { $size ->
       [1572864000] 1,5GB
       *[other] 50 MB
   }, which is my maximum limit.
downloading = Downloading...
uploading = Uploading...
youtube-error =
    <b>An error occurred while processing the video, try again later.</b>
    If the problem persists, please contact my developer.
auto-help = When enabled, the bot will automatically download photos and videos from certain social networks upon detecting a link, removing the need for the /sdl or /dl command.
caption-help = When enabled, the caption of medias downloaded via the bot will be sent along with the media.
no-link-provided =
    <b>No link provided or the link is invalid.</b>
    You need to specify a valid link from <b><i>Instagram</i></b>, <b><i>TikTok</i></b>, <b><i>Reddit</i></b>, <b><i>Twitter/X</i></b>, <b><i>Threads</i></b>, <b><i>Reddit</i></b>, <b><i>BlueSky</i></b>, or <b><i>YouTube Shorts</i></b> to download the media.
misc = Misc
misc-help =
    <b>Miscellaneous</b>
    
    This module combines some useful commands that don't fit into other specific categories.
    
    <b>— Commands:</b>
    <b>/weather (city):</b> Displays the current weather of the specified city.
    <b>/tr (source)-(destination) (text):</b> Translates a text from the source language to the specified destination language.
    <i>If you don't specify the source language, Smudge will automatically detect it.</i>

    <b>Note:</b>
    You can translate messages by replying to them with <code>/translate</code>.
    Both <code>/tr</code> and <code>/translate</code> commands work the same way.
translator-no-args-provided =
    You need to specify the text you want to translate or reply to a text message, or a photo with a caption.

    <b>Usage:</b> <code>/tr (?language) (text for translation)</code>
weather-no-location-provided =
    You need to specify the location for which you want to know the weather information.
    
    <b>Example:</b> <code>/weather Belém</code>.
weather-select-location = <b>Select the location where you want to know the Weather:</b>
weather-details =
    <b>{ $localname }</b>:

    Temperature: <code>{ $temperature } °C</code>
    Temperature feels like: <code>{ $temperatureFeelsLike } °C</code>
    Air humidity: <code>{ $relativeHumidity }%</code>
    Wind speed: <code>{ $windSpeed } km/h</code>
stickers = Stickers
kanging = <code>Kanging (Stealing) the sticker...</code>
kang-no-reply-provided = You need to use this command by replying to <i><b>a sticker</b></i>, <i><b>an image</b></i> or <i><b>a video</b></i>.
converting-video-to-sticker = <code>Converting video/gif to video sticker...</code>
sticker-pack-already-exists = <code>Using existing sticker pack...</code>
kang-error =
    <b>An error occurred while processing the sticker, try again later.</b>
    If the problem persists, please contact my developer.
get-sticker-no-reply-provided =
    You need to use this command by replying to an <b>static (image) or video sticker.
sticker-invalid-media-type = The file you replied to is not valid. You need to reply to an <i><b>sticker</b></i>, <i><b>video</b></i> or <i><b>photo</b></i>.
sticker-new-pack = <code>Creating a new sticker pack...</code>
sticker-stoled = 
    Sticker <b>successfully</b> stolen, <a href='t.me/addstickers/{ $stickerSetName }'>check out.</a>
    <b>Emoji:</b> { $emoji }
stickers-help = 
    <b>Stickers</b>

    This module contains some useful functions for you to manage stickers.

    <b>— Commands:</b>
    <b>/kang (emoji):</b> Reply to any sticker to add it to your sticker pack created by me. <b>Works with <i>static, video, and animated</i> stickers.</b>
    <b>/getsticker:</b> Reply to a sticker to be able to send it as a <i>.png</i> or <i>.gif</i> file. <b>Only works with <i>static or animated</i> stickers.</b>
lastfm = Last.FM
no-lastfm-username-provided =
    You need to specify your last.fm username so that I can save my database.
    
    <b>Example:</b> <code>/setuser maozedong</code>.
invalid-lastfm-username =
    <b>Invalid last.fm username</b>
    Check that you have correctly typed your last.FM username and try again.
lastfm-username-not-defined =
    <b>You haven't set your last.fm username yet.</b>
    Use the command /setuser to set.
lastfm-username-saved = <b>Done</b>, your last.fm username has been saved!
lastfm-error =
    <b>An error seems to have occurred.</b>
    Last.fm might be temporarily unavailable.

    Please try again later. If the problem persists, please contact my developer
no-scrobbled-yet = 
    <b>It looks like you haven't scrobbled any tracks on Last.fm yet.</b>

    If you're facing issues with Last.fm, visit last.fm/about/trackmymusic to learn how to connect your account to your music app.
lastfm-playing = 
    <b><a href='https://last.fm/user/{ $lastFMUsername }'>{ $firstName }</a></b> { $nowplaying ->
       [true] is listening for the 
      *[false] was listening for the 
   } { NUMBER($playcount, type: "ordinal") ->
       [1] { $playcount }st
       [2] { $playcount }nd
       [3] { $playcount }rd
       *[other] { $playcount }th
   } time:
lastfm-help =
    <b>Last.FM Scobbles</b>

    <b>Scrobble</b> is a feature of last.fm that automatically sends information about the music you're listening to to the service.
    <b>To know more <a href='https://www.last.fm/pt/about/trackmymusic'>click here</a></b>.

    <b>— Commands:</b>
    <b>/setuser (lastfm username):</b> Set your last.fm username.
    <b>/lastfm | /lp:</b> Shows the music you are or were listening to.
    <b>/album | /alb:</b>Shows the album you are or were listening to.
    <b>/artist   | /art:</b> Shows the artist you are or were listening to.