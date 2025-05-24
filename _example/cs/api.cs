




/**
 * 广告Banner接口
 */
public class AdBannerApi
{

    private IApiClient client;

    public AdBannerApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 
	 */
    public Task<ApiResultListAdBannerViewEntity> list() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultListAdBannerViewEntity>("/pb/ad-banner/list",param);
    }

	
}



/**
 * 广告兴趣社接口
 */
public class AdInterestApi
{

    private IApiClient client;

    public AdInterestApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 数据
	 */
    public Task<ApiResultAdInterestResult> dataset() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultAdInterestResult>("/pt/ad-interest/dataset",param);
    }

	
}



/**
 * 广告推广接口
 */
public class AdPromotionApi
{

    private IApiClient client;

    public AdPromotionApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 隔日唤起-用户跳转完成回调
	 */
    public Task<ApiResultBoolean> jumpCompleted( IdentityPairLongString body) {
        
        return client.Put<,>("/pt/ad-promotion/launch/completed",body);
    }

	
	/**
	 * 
	 * 隔日唤起-获取当日跳转信息
	 */
    public Task<ApiResultListAdNextDayLaunchResult> launch() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultListAdNextDayLaunchResult>("/pt/ad-promotion/launch",param);
    }

	
	/**
	 * 
	 * 获取当前推广进度
	 */
    public Task<ApiResultGlobalPromotionProgress> progress() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultGlobalPromotionProgress>("/pt/ad-promotion/progress",param);
    }

	
	/**
	 * 
	 * 查询备选任务
	 */
    public Task<ApiResultListAdPromotionJobEntity> slaves( long id) {
        var param = new Dictionary<string,dynamic>();
        param["id"] = id;
        
        return client.Get<ApiResultListAdPromotionJobEntity>("/pt/ad-promotion/slaves",param);
    }

	
	/**
	 * 
	 * 提交任务
	 */
    public Task<ApiResultGlobalPromotionFinish> submit_1( GlobalPromotionForm body) {
        
        return client.Post<,>("/pt/ad-promotion/submit",body);
    }

	
	/**
	 * 
	 * 累计获取奖励
	 */
    public Task<ApiResultRewardTip> totalRewardTip() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultRewardTip>("/pt/ad-promotion/total-reward-tip",param);
    }

	
}



/**
 * 应用升级接口
 */
public class AppUpgradeApi
{

    private IApiClient client;

    public AppUpgradeApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 校验
	 */
    public Task<ApiResultUpgradeResult> check( string platform, string version) {
        var param = new Dictionary<string,dynamic>();
        param["platform"] = platform;
        param["version"] = version;
        
        return client.Get<ApiResultUpgradeResult>("/pb/upgrade/check",param);
    }

	
}



/**
 * 语音包接口
 */
public class AudioPackageApi
{

    private IApiClient client;

    public AudioPackageApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 添加行为
	 */
    public Task<ApiResultBoolean> addInteractive( UserBehaviorArg body) {
        
        return client.Post<,>("/pt/audio-package/interactive",body);
    }

	
	/**
	 * 
	 * 标准分类查询
	 */
    public Task<ApiResultListIdentityPairStringString> category() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultListIdentityPairStringString>("/pt/audio-package/category",param);
    }

	
	/**
	 * 
	 * 语音包详情
	 */
    public Task<ApiResultAudioPackageViewEntry> detail_1( long packageId) {
        var param = new Dictionary<string,dynamic>();
        param["packageId"] = packageId;
        
        return client.Get<ApiResultAudioPackageViewEntry>("/pt/audio-package/detail",param);
    }

	
	/**
	 * 
	 * 音频集播放参数
	 */
    public Task<ApiResultMapStringString> files_1() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultMapStringString>("/pt/audio-package/player-config",param);
    }

	
	/**
	 * 
	 * 音频集
	 */
    public Task<ApiResultListAudioFileEntityViewEntry> files_2( long packageId) {
        var param = new Dictionary<string,dynamic>();
        param["packageId"] = packageId;
        
        return client.Get<ApiResultListAudioFileEntityViewEntry>("/pt/audio-package/files",param);
    }

	
	/**
	 * 
	 * 查询行为
	 */
    public Task<ApiResultUserBehaviorReference> getInteractive( long packageId) {
        var param = new Dictionary<string,dynamic>();
        param["packageId"] = packageId;
        
        return client.Get<ApiResultUserBehaviorReference>("/pt/audio-package/interactive",param);
    }

	
	/**
	 * 
	 * 统计数据
	 */
    public Task<ApiResultStatisticReference> getStatistic( long packageId) {
        var param = new Dictionary<string,dynamic>();
        param["packageId"] = packageId;
        
        return client.Get<ApiResultStatisticReference>("/pt/audio-package/statistic",param);
    }

	
	/**
	 * 
	 * 隐藏音频集
	 */
    public Task<ApiResultListAudioFileEntity> hiddenFiles( long packageId) {
        var param = new Dictionary<string,dynamic>();
        param["packageId"] = packageId;
        
        return client.Get<ApiResultListAudioFileEntity>("/pt/audio-package/hidden-files",param);
    }

	
	/**
	 * 
	 * 分类列表查询
	 */
    public Task<ApiResultPageDataAudioPackageEntity> page_2( string categoryId, string name, int page, int size) {
        var param = new Dictionary<string,dynamic>();
        param["categoryId"] = categoryId;
        param["name"] = name;
        param["page"] = page;
        param["size"] = size;
        
        return client.Get<ApiResultPageDataAudioPackageEntity>("/pt/audio-package/page",param);
    }

	
	/**
	 * 
	 * 推荐标签查询
	 */
    public Task<ApiResultListKVPairStringString> recommendTags() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultListKVPairStringString>("/pt/audio-package/recommend-tags",param);
    }

	
	/**
	 * 
	 * 推荐列表查询
	 */
    public Task<ApiResultListAudioPackageEntity> recommend_1( string type) {
        var param = new Dictionary<string,dynamic>();
        param["type"] = type;
        
        return client.Get<ApiResultListAudioPackageEntity>("/pt/audio-package/recommend",param);
    }

	
	/**
	 * 
	 * 删除行为
	 */
    public Task<ApiResultBoolean> removeInteractive( UserBehaviorArg body) {
        
        return client.Delete<,>("/pt/audio-package/interactive",body);
    }

	
}



/**
 * 授权登陆
 */
public class AuthorizationApi
{

    private IApiClient client;

    public AuthorizationApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 激活设备
	 */
    public Task<ApiResultBoolean> activation( DeviceActivation body) {
        
        return client.Post<,>("/pm/authorization/activation",body);
    }

	
	/**
	 * 
	 * 获取验证码
	 */
    public Task<ApiResultBoolean> getCode( string phone, string zone) {
        var param = new Dictionary<string,dynamic>();
        param["phone"] = phone;
        param["zone"] = zone;
        
        return client.Post<,>("/pm/authorization/send_code",param);
    }

	
}



/**
 * 提现订单
 */
public class CashOrderApi
{

    private IApiClient client;

    public CashOrderApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 申请
	 */
    public Task<ApiResultLong> apply( CashOrderCreate body) {
        
        return client.Post<,>("/pt/cash/order/apply",body);
    }

	
	/**
	 * 
	 * 明细
	 */
    public Task<ApiResultCashOrderEntity> detail( long id) {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultCashOrderEntity>(string.Format("/pt/cash/order/detail/{0}",id),param);
    }

	
}



/**
 * OSS对象签名
 */
public class GenericOssPresignProvider
{

    private IApiClient client;

    public GenericOssPresignProvider(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * ACL授权
	 * 
	 */
    public Task<ApiResultListKVPairStringString> presign( List<string> body) {
        
        return client.Post<,>("/pt/storage/oss/presign",body);
    }

	
	/**
	 * 生成上传路径
	 * 
	 */
    public Task<ApiResultKVPairStringString> uploadUrl( FileLocation body) {
        
        return client.Post<,>("/pt/storage/oss/presign/upload-url",body);
    }

	
}



/**
 * 产品接口
 */
public class ProductApi
{

    private IApiClient client;

    public ProductApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 推广兑换会员产品
	 */
    public Task<ApiResultVipProductEntity> exchangeVip() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultVipProductEntity>("/pb/product/vips/exchange/ad",param);
    }

	
	/**
	 * 
	 * 续费会员产品集
	 */
    public Task<ApiResultListVipProductEntity> renewalVips() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultListVipProductEntity>("/pb/product/vips/renewal",param);
    }

	
	/**
	 * 
	 * 实时变声推荐
	 */
    public Task<ApiResultListTimbreProductEntity> stsRecommend( string type) {
        var param = new Dictionary<string,dynamic>();
        param["type"] = type;
        
        return client.Get<ApiResultListTimbreProductEntity>("/pb/product/sts/recommend",param);
    }

	
	/**
	 * 
	 * 实时变声音色集
	 */
    public Task<ApiResultListTimbreProductEntity> stsTimbres() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultListTimbreProductEntity>("/pb/product/sts/timbres",param);
    }

	
	/**
	 * 
	 * 文字变声推荐
	 */
    public Task<ApiResultListTimbreProductEntity> ttsRecommend( string type) {
        var param = new Dictionary<string,dynamic>();
        param["type"] = type;
        
        return client.Get<ApiResultListTimbreProductEntity>("/pb/product/tts/recommend",param);
    }

	
	/**
	 * 
	 * 文转音音色集
	 */
    public Task<ApiResultListTimbreProductEntity> ttsTimbres() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultListTimbreProductEntity>("/pb/product/tts/timbres",param);
    }

	
	/**
	 * 
	 * 充值会员产品集
	 */
    public Task<ApiResultListVipProductEntity> vips() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultListVipProductEntity>("/pb/product/vips",param);
    }

	
}



/**
 * 充值折扣接口
 */
public class RechargeDiscountApi
{

    private IApiClient client;

    public RechargeDiscountApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 订单优惠券预览
	 */
    public Task<ApiResultListUserCouponEntity> coupons( GeneralOrderDiscountPreview body) {
        
        return client.Post<,>("/pt/recharge/discount/coupons",body);
    }

	
	/**
	 * 
	 * 订单折扣预览
	 */
    public Task<ApiResultGeneralOrderDiscountPreviewResult> preview( GeneralOrderDiscountPreview body) {
        
        return client.Post<,>("/pt/recharge/discount/preview",body);
    }

	
	/**
	 * 
	 * 订单折扣推荐
	 */
    public Task<ApiResultGeneralOrderDiscountPreviewResult> recommend( GeneralOrderDiscountPreview body) {
        
        return client.Post<,>("/pt/recharge/discount/recommend",body);
    }

	
}



/**
 * 充值订单接口
 */
public class RechargeOrderApi
{

    private IApiClient client;

    public RechargeOrderApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 取消订单
	 */
    public Task<ApiResultBoolean> cancel( GeneralOrderCancel body) {
        
        return client.Put<,>("/pt/recharge/order/cancel",body);
    }

	
	/**
	 * 
	 * 创建订单
	 */
    public Task<ApiResultGeneralOrderCreateResult> create( GeneralOrderCreate body) {
        
        return client.Post<,>("/pt/recharge/order/create",body);
    }

	
	/**
	 * 
	 * 查询订单状态
	 */
    public Task<ApiResultGeneralOrderState> state_1( long id) {
        var param = new Dictionary<string,dynamic>();
        param["id"] = id;
        
        return client.Get<ApiResultGeneralOrderState>("/pt/recharge/order/state",param);
    }

	
}



/**
 * 举报接口
 */
public class ReportContentApi
{

    private IApiClient client;

    public ReportContentApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 举报
	 */
    public Task<ApiResultLong> submit( ReportContentEntity body) {
        
        return client.Post<,>("/pt/report/submit",body);
    }

	
}



/**
 * 用户素材库接口
 */
public class UserAudioPackageApi
{

    private IApiClient client;

    public UserAudioPackageApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 提交音频
	 */
    public Task<ApiResultLong> addFile( KVPairStringString body) {
        
        return client.Post<,>("/pt/user-audio-package/file/add",body);
    }

	
	/**
	 * 
	 * 查询审核结果
	 */
    public Task<ApiResultAudioPackageAuditResult> auditQuery( long packageId) {
        var param = new Dictionary<string,dynamic>();
        param["packageId"] = packageId;
        
        return client.Get<ApiResultAudioPackageAuditResult>("/pt/user-audio-package/audit/result",param);
    }

	
	/**
	 * 
	 * 删除音频
	 */
    public Task<ApiResultBoolean> deleteFile( IdentityPairLongLong body) {
        
        return client.Delete<,>("/pt/user-audio-package/file/delete",body);
    }

	
	/**
	 * 
	 * 查询音频列表
	 */
    public Task<ApiResultPageDataAudioFileEntity> files( int page, int size) {
        var param = new Dictionary<string,dynamic>();
        param["page"] = page;
        param["size"] = size;
        
        return client.Get<ApiResultPageDataAudioFileEntity>("/pt/user-audio-package/files",param);
    }

	
	/**
	 * 
	 * 生成语音包ID
	 */
    public Task<ApiResultLong> newPackage() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultLong>("/pt/user-audio-package/new",param);
    }

	
	/**
	 * 
	 * 语音包列表
	 */
    public Task<ApiResultPageDataAudioPackageEntity> page_1( int page, int size) {
        var param = new Dictionary<string,dynamic>();
        param["page"] = page;
        param["size"] = size;
        
        return client.Get<ApiResultPageDataAudioPackageEntity>("/pt/user-audio-package/page",param);
    }

	
	/**
	 * 
	 * 提交
	 */
    public Task<ApiResultLong> submitPackage( UserAudioPackageUploadForm body) {
        
        return client.Post<,>("/pt/user-audio-package/submit",body);
    }

	
	/**
	 * 
	 * 语音包优惠券兑换
	 */
    public Task<ApiResultBoolean> unlock( AudioCouponUnlockForm body) {
        
        return client.Post<,>("/pt/user-audio-package/unlock",body);
    }

	
}



/**
 * 用户账户接口
 */
public class UserFinanceApi
{

    private IApiClient client;

    public UserFinanceApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 提现明细
	 */
    public Task<ApiResultPageDataCashOrderEntity> cash( int page, int size, List<int> state) {
        var param = new Dictionary<string,dynamic>();
        param["page"] = page;
        param["size"] = size;
        param["state"] = state;
        
        return client.Get<ApiResultPageDataCashOrderEntity>("/pt/user/finance/cash/page",param);
    }

	
	/**
	 * 
	 * 用户账户
	 */
    public Task<ApiResultUserFinanceAccount> finance() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultUserFinanceAccount>("/pt/user/finance",param);
    }

	
	/**
	 * 
	 * 收益明细
	 */
    public Task<ApiResultPageDataUserFinanceEntryDetail> profit( int page, int size, int state) {
        var param = new Dictionary<string,dynamic>();
        param["page"] = page;
        param["size"] = size;
        param["state"] = state;
        
        return client.Get<ApiResultPageDataUserFinanceEntryDetail>("/pt/user/finance/profit/page",param);
    }

	
}



/**
 * 用户画像接口
 */
public class UserProfileApi
{

    private IApiClient client;

    public UserProfileApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 
	 * 注销
	 */
    public Task<ApiResultBoolean> delete( IdentityPairLongString body) {
        
        return client.Put<,>("/pt/user/destroy",body);
    }

	
	/**
	 * 
	 * 初始化用户信息
	 */
    public Task<ApiResultBoolean> init( UserInitProfile body) {
        
        return client.Put<,>("/pt/user/profile/init",body);
    }

	
	/**
	 * 
	 * 我的收藏的音频
	 */
    public Task<ApiResultPageDataAudioFileEntity> myFavorite( int page, int size) {
        var param = new Dictionary<string,dynamic>();
        param["page"] = page;
        param["size"] = size;
        
        return client.Get<ApiResultPageDataAudioFileEntity>("/pt/user/profile/favorite/audio",param);
    }

	
	/**
	 * 
	 * 我的邀请记录
	 */
    public Task<ApiResultPageDataUserInviteLog> myInvited( int page, int size) {
        var param = new Dictionary<string,dynamic>();
        param["page"] = page;
        param["size"] = size;
        
        return client.Get<ApiResultPageDataUserInviteLog>("/pt/user/profile/invited",param);
    }

	
	/**
	 * 
	 * 我的会员
	 */
    public Task<ApiResultUserVipRechargeAccount> myVip() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultUserVipRechargeAccount>("/pt/user/profile/vip",param);
    }

	
	/**
	 * 
	 * 查询用户信息
	 */
    public Task<ApiResultUserPrincipalEntity> profile() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultUserPrincipalEntity>("/pt/user/profile",param);
    }

	
	/**
	 * 
	 * 更新用户信息
	 */
    public Task<ApiResultBoolean> update( UserInitProfile body) {
        
        return client.Put<,>("/pt/user/profile/update",body);
    }

	
}











