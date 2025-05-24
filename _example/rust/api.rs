

use std::collections::HashMap;
use std::io::Error;
use std::iter::Map;
use crate::Client::ApiClient;
use crate::Model::*;




/**
 * 广告Banner接口
 */
pub struct AdBannerApi {
    client: dyn ApiClient,
}

impl AdBannerApi {

	

	/**
	 * 
	 * 
	 */
    async fn list(&self, ) -> Result<ApiResultListAdBannerViewEntity,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pb/ad-banner/list"),params)
    }

	
}



/**
 * 广告兴趣社接口
 */
pub struct AdInterestApi {
    client: dyn ApiClient,
}

impl AdInterestApi {

	

	/**
	 * 
	 * 数据
	 */
    async fn dataset(&self, ) -> Result<ApiResultAdInterestResult,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/ad-interest/dataset"),params)
    }

	
}



/**
 * 广告推广接口
 */
pub struct AdPromotionApi {
    client: dyn ApiClient,
}

impl AdPromotionApi {

	

	/**
	 * 
	 * 隔日唤起-用户跳转完成回调
	 */
    async fn jumpCompleted(&self, body: IdentityPairLongString) -> Result<ApiResultBoolean,Error>  {
        
        self.client.put(String::from("/pt/ad-promotion/launch/completed"),body)
    }

	

	/**
	 * 
	 * 隔日唤起-获取当日跳转信息
	 */
    async fn launch(&self, ) -> Result<ApiResultListAdNextDayLaunchResult,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/ad-promotion/launch"),params)
    }

	

	/**
	 * 
	 * 获取当前推广进度
	 */
    async fn progress(&self, ) -> Result<ApiResultGlobalPromotionProgress,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/ad-promotion/progress"),params)
    }

	

	/**
	 * 
	 * 查询备选任务
	 */
    async fn slaves(&self, id: u64) -> Result<ApiResultListAdPromotionJobEntity,Error>  {
        let mut params = HashMap::new();
        params.insert("id",id.to_string());
        
        self.client.get(String::from("/pt/ad-promotion/slaves"),params)
    }

	

	/**
	 * 
	 * 提交任务
	 */
    async fn submit_1(&self, body: GlobalPromotionForm) -> Result<ApiResultGlobalPromotionFinish,Error>  {
        
        self.client.post(String::from("/pt/ad-promotion/submit"),body)
    }

	

	/**
	 * 
	 * 累计获取奖励
	 */
    async fn totalRewardTip(&self, ) -> Result<ApiResultRewardTip,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/ad-promotion/total-reward-tip"),params)
    }

	
}



/**
 * 应用升级接口
 */
pub struct AppUpgradeApi {
    client: dyn ApiClient,
}

impl AppUpgradeApi {

	

	/**
	 * 
	 * 校验
	 */
    async fn check(&self, platform: String,version: String) -> Result<ApiResultUpgradeResult,Error>  {
        let mut params = HashMap::new();
        params.insert("platform",platform.to_string());
        params.insert("version",version.to_string());
        
        self.client.get(String::from("/pb/upgrade/check"),params)
    }

	
}



/**
 * 语音包接口
 */
pub struct AudioPackageApi {
    client: dyn ApiClient,
}

impl AudioPackageApi {

	

	/**
	 * 
	 * 添加行为
	 */
    async fn addInteractive(&self, body: UserBehaviorArg) -> Result<ApiResultBoolean,Error>  {
        
        self.client.post(String::from("/pt/audio-package/interactive"),body)
    }

	

	/**
	 * 
	 * 标准分类查询
	 */
    async fn category(&self, ) -> Result<ApiResultListIdentityPairStringString,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/audio-package/category"),params)
    }

	

	/**
	 * 
	 * 语音包详情
	 */
    async fn detail_1(&self, packageId: u64) -> Result<ApiResultAudioPackageViewEntry,Error>  {
        let mut params = HashMap::new();
        params.insert("packageId",packageId.to_string());
        
        self.client.get(String::from("/pt/audio-package/detail"),params)
    }

	

	/**
	 * 
	 * 音频集播放参数
	 */
    async fn files_1(&self, ) -> Result<ApiResultMapStringString,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/audio-package/player-config"),params)
    }

	

	/**
	 * 
	 * 音频集
	 */
    async fn files_2(&self, packageId: u64) -> Result<ApiResultListAudioFileEntityViewEntry,Error>  {
        let mut params = HashMap::new();
        params.insert("packageId",packageId.to_string());
        
        self.client.get(String::from("/pt/audio-package/files"),params)
    }

	

	/**
	 * 
	 * 查询行为
	 */
    async fn getInteractive(&self, packageId: u64) -> Result<ApiResultUserBehaviorReference,Error>  {
        let mut params = HashMap::new();
        params.insert("packageId",packageId.to_string());
        
        self.client.get(String::from("/pt/audio-package/interactive"),params)
    }

	

	/**
	 * 
	 * 统计数据
	 */
    async fn getStatistic(&self, packageId: u64) -> Result<ApiResultStatisticReference,Error>  {
        let mut params = HashMap::new();
        params.insert("packageId",packageId.to_string());
        
        self.client.get(String::from("/pt/audio-package/statistic"),params)
    }

	

	/**
	 * 
	 * 隐藏音频集
	 */
    async fn hiddenFiles(&self, packageId: u64) -> Result<ApiResultListAudioFileEntity,Error>  {
        let mut params = HashMap::new();
        params.insert("packageId",packageId.to_string());
        
        self.client.get(String::from("/pt/audio-package/hidden-files"),params)
    }

	

	/**
	 * 
	 * 分类列表查询
	 */
    async fn page_2(&self, categoryId: String,name: String,page: usize,size: usize) -> Result<ApiResultPageDataAudioPackageEntity,Error>  {
        let mut params = HashMap::new();
        params.insert("categoryId",categoryId.to_string());
        params.insert("name",name.to_string());
        params.insert("page",page.to_string());
        params.insert("size",size.to_string());
        
        self.client.get(String::from("/pt/audio-package/page"),params)
    }

	

	/**
	 * 
	 * 推荐标签查询
	 */
    async fn recommendTags(&self, ) -> Result<ApiResultListKVPairStringString,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/audio-package/recommend-tags"),params)
    }

	

	/**
	 * 
	 * 推荐列表查询
	 */
    async fn recommend_1(&self, type: String) -> Result<ApiResultListAudioPackageEntity,Error>  {
        let mut params = HashMap::new();
        params.insert("type",type.to_string());
        
        self.client.get(String::from("/pt/audio-package/recommend"),params)
    }

	

	/**
	 * 
	 * 删除行为
	 */
    async fn removeInteractive(&self, body: UserBehaviorArg) -> Result<ApiResultBoolean,Error>  {
        
        self.client.delete(String::from("/pt/audio-package/interactive"),body)
    }

	
}



/**
 * 授权登陆
 */
pub struct AuthorizationApi {
    client: dyn ApiClient,
}

impl AuthorizationApi {

	

	/**
	 * 
	 * 激活设备
	 */
    async fn activation(&self, body: DeviceActivation) -> Result<ApiResultBoolean,Error>  {
        
        self.client.post(String::from("/pm/authorization/activation"),body)
    }

	

	/**
	 * 
	 * 获取验证码
	 */
    async fn getCode(&self, phone: String,zone: String) -> Result<ApiResultBoolean,Error>  {
        let mut params = HashMap::new();
        params.insert("phone",phone.to_string());
        params.insert("zone",zone.to_string());
        
        self.client.post(String::from("/pm/authorization/send_code"),params)
    }

	
}



/**
 * 提现订单
 */
pub struct CashOrderApi {
    client: dyn ApiClient,
}

impl CashOrderApi {

	

	/**
	 * 
	 * 申请
	 */
    async fn apply(&self, body: CashOrderCreate) -> Result<ApiResultLong,Error>  {
        
        self.client.post(String::from("/pt/cash/order/apply"),body)
    }

	

	/**
	 * 
	 * 明细
	 */
    async fn detail(&self, id: u64) -> Result<ApiResultCashOrderEntity,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from(format!("/pt/cash/order/detail/${id}",id)),params)
    }

	
}



/**
 * OSS对象签名
 */
pub struct GenericOssPresignProvider {
    client: dyn ApiClient,
}

impl GenericOssPresignProvider {

	

	/**
	 * ACL授权
	 * 
	 */
    async fn presign(&self, body: Vec<String>) -> Result<ApiResultListKVPairStringString,Error>  {
        
        self.client.post(String::from("/pt/storage/oss/presign"),body)
    }

	

	/**
	 * 生成上传路径
	 * 
	 */
    async fn uploadUrl(&self, body: FileLocation) -> Result<ApiResultKVPairStringString,Error>  {
        
        self.client.post(String::from("/pt/storage/oss/presign/upload-url"),body)
    }

	
}



/**
 * 产品接口
 */
pub struct ProductApi {
    client: dyn ApiClient,
}

impl ProductApi {

	

	/**
	 * 
	 * 推广兑换会员产品
	 */
    async fn exchangeVip(&self, ) -> Result<ApiResultVipProductEntity,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pb/product/vips/exchange/ad"),params)
    }

	

	/**
	 * 
	 * 续费会员产品集
	 */
    async fn renewalVips(&self, ) -> Result<ApiResultListVipProductEntity,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pb/product/vips/renewal"),params)
    }

	

	/**
	 * 
	 * 实时变声推荐
	 */
    async fn stsRecommend(&self, type: String) -> Result<ApiResultListTimbreProductEntity,Error>  {
        let mut params = HashMap::new();
        params.insert("type",type.to_string());
        
        self.client.get(String::from("/pb/product/sts/recommend"),params)
    }

	

	/**
	 * 
	 * 实时变声音色集
	 */
    async fn stsTimbres(&self, ) -> Result<ApiResultListTimbreProductEntity,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pb/product/sts/timbres"),params)
    }

	

	/**
	 * 
	 * 文字变声推荐
	 */
    async fn ttsRecommend(&self, type: String) -> Result<ApiResultListTimbreProductEntity,Error>  {
        let mut params = HashMap::new();
        params.insert("type",type.to_string());
        
        self.client.get(String::from("/pb/product/tts/recommend"),params)
    }

	

	/**
	 * 
	 * 文转音音色集
	 */
    async fn ttsTimbres(&self, ) -> Result<ApiResultListTimbreProductEntity,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pb/product/tts/timbres"),params)
    }

	

	/**
	 * 
	 * 充值会员产品集
	 */
    async fn vips(&self, ) -> Result<ApiResultListVipProductEntity,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pb/product/vips"),params)
    }

	
}



/**
 * 充值折扣接口
 */
pub struct RechargeDiscountApi {
    client: dyn ApiClient,
}

impl RechargeDiscountApi {

	

	/**
	 * 
	 * 订单优惠券预览
	 */
    async fn coupons(&self, body: GeneralOrderDiscountPreview) -> Result<ApiResultListUserCouponEntity,Error>  {
        
        self.client.post(String::from("/pt/recharge/discount/coupons"),body)
    }

	

	/**
	 * 
	 * 订单折扣预览
	 */
    async fn preview(&self, body: GeneralOrderDiscountPreview) -> Result<ApiResultGeneralOrderDiscountPreviewResult,Error>  {
        
        self.client.post(String::from("/pt/recharge/discount/preview"),body)
    }

	

	/**
	 * 
	 * 订单折扣推荐
	 */
    async fn recommend(&self, body: GeneralOrderDiscountPreview) -> Result<ApiResultGeneralOrderDiscountPreviewResult,Error>  {
        
        self.client.post(String::from("/pt/recharge/discount/recommend"),body)
    }

	
}



/**
 * 充值订单接口
 */
pub struct RechargeOrderApi {
    client: dyn ApiClient,
}

impl RechargeOrderApi {

	

	/**
	 * 
	 * 取消订单
	 */
    async fn cancel(&self, body: GeneralOrderCancel) -> Result<ApiResultBoolean,Error>  {
        
        self.client.put(String::from("/pt/recharge/order/cancel"),body)
    }

	

	/**
	 * 
	 * 创建订单
	 */
    async fn create(&self, body: GeneralOrderCreate) -> Result<ApiResultGeneralOrderCreateResult,Error>  {
        
        self.client.post(String::from("/pt/recharge/order/create"),body)
    }

	

	/**
	 * 
	 * 查询订单状态
	 */
    async fn state_1(&self, id: u64) -> Result<ApiResultGeneralOrderState,Error>  {
        let mut params = HashMap::new();
        params.insert("id",id.to_string());
        
        self.client.get(String::from("/pt/recharge/order/state"),params)
    }

	
}



/**
 * 举报接口
 */
pub struct ReportContentApi {
    client: dyn ApiClient,
}

impl ReportContentApi {

	

	/**
	 * 
	 * 举报
	 */
    async fn submit(&self, body: ReportContentEntity) -> Result<ApiResultLong,Error>  {
        
        self.client.post(String::from("/pt/report/submit"),body)
    }

	
}



/**
 * 用户素材库接口
 */
pub struct UserAudioPackageApi {
    client: dyn ApiClient,
}

impl UserAudioPackageApi {

	

	/**
	 * 
	 * 提交音频
	 */
    async fn addFile(&self, body: KVPairStringString) -> Result<ApiResultLong,Error>  {
        
        self.client.post(String::from("/pt/user-audio-package/file/add"),body)
    }

	

	/**
	 * 
	 * 查询审核结果
	 */
    async fn auditQuery(&self, packageId: u64) -> Result<ApiResultAudioPackageAuditResult,Error>  {
        let mut params = HashMap::new();
        params.insert("packageId",packageId.to_string());
        
        self.client.get(String::from("/pt/user-audio-package/audit/result"),params)
    }

	

	/**
	 * 
	 * 删除音频
	 */
    async fn deleteFile(&self, body: IdentityPairLongLong) -> Result<ApiResultBoolean,Error>  {
        
        self.client.delete(String::from("/pt/user-audio-package/file/delete"),body)
    }

	

	/**
	 * 
	 * 查询音频列表
	 */
    async fn files(&self, page: usize,size: usize) -> Result<ApiResultPageDataAudioFileEntity,Error>  {
        let mut params = HashMap::new();
        params.insert("page",page.to_string());
        params.insert("size",size.to_string());
        
        self.client.get(String::from("/pt/user-audio-package/files"),params)
    }

	

	/**
	 * 
	 * 生成语音包ID
	 */
    async fn newPackage(&self, ) -> Result<ApiResultLong,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/user-audio-package/new"),params)
    }

	

	/**
	 * 
	 * 语音包列表
	 */
    async fn page_1(&self, page: usize,size: usize) -> Result<ApiResultPageDataAudioPackageEntity,Error>  {
        let mut params = HashMap::new();
        params.insert("page",page.to_string());
        params.insert("size",size.to_string());
        
        self.client.get(String::from("/pt/user-audio-package/page"),params)
    }

	

	/**
	 * 
	 * 提交
	 */
    async fn submitPackage(&self, body: UserAudioPackageUploadForm) -> Result<ApiResultLong,Error>  {
        
        self.client.post(String::from("/pt/user-audio-package/submit"),body)
    }

	

	/**
	 * 
	 * 语音包优惠券兑换
	 */
    async fn unlock(&self, body: AudioCouponUnlockForm) -> Result<ApiResultBoolean,Error>  {
        
        self.client.post(String::from("/pt/user-audio-package/unlock"),body)
    }

	
}



/**
 * 用户账户接口
 */
pub struct UserFinanceApi {
    client: dyn ApiClient,
}

impl UserFinanceApi {

	

	/**
	 * 
	 * 提现明细
	 */
    async fn cash(&self, page: usize,size: usize,state: Vec<usize>) -> Result<ApiResultPageDataCashOrderEntity,Error>  {
        let mut params = HashMap::new();
        params.insert("page",page.to_string());
        params.insert("size",size.to_string());
        params.insert("state",state.to_string());
        
        self.client.get(String::from("/pt/user/finance/cash/page"),params)
    }

	

	/**
	 * 
	 * 用户账户
	 */
    async fn finance(&self, ) -> Result<ApiResultUserFinanceAccount,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/user/finance"),params)
    }

	

	/**
	 * 
	 * 收益明细
	 */
    async fn profit(&self, page: usize,size: usize,state: usize) -> Result<ApiResultPageDataUserFinanceEntryDetail,Error>  {
        let mut params = HashMap::new();
        params.insert("page",page.to_string());
        params.insert("size",size.to_string());
        params.insert("state",state.to_string());
        
        self.client.get(String::from("/pt/user/finance/profit/page"),params)
    }

	
}



/**
 * 用户画像接口
 */
pub struct UserProfileApi {
    client: dyn ApiClient,
}

impl UserProfileApi {

	

	/**
	 * 
	 * 注销
	 */
    async fn delete(&self, body: IdentityPairLongString) -> Result<ApiResultBoolean,Error>  {
        
        self.client.put(String::from("/pt/user/destroy"),body)
    }

	

	/**
	 * 
	 * 初始化用户信息
	 */
    async fn init(&self, body: UserInitProfile) -> Result<ApiResultBoolean,Error>  {
        
        self.client.put(String::from("/pt/user/profile/init"),body)
    }

	

	/**
	 * 
	 * 我的收藏的音频
	 */
    async fn myFavorite(&self, page: usize,size: usize) -> Result<ApiResultPageDataAudioFileEntity,Error>  {
        let mut params = HashMap::new();
        params.insert("page",page.to_string());
        params.insert("size",size.to_string());
        
        self.client.get(String::from("/pt/user/profile/favorite/audio"),params)
    }

	

	/**
	 * 
	 * 我的邀请记录
	 */
    async fn myInvited(&self, page: usize,size: usize) -> Result<ApiResultPageDataUserInviteLog,Error>  {
        let mut params = HashMap::new();
        params.insert("page",page.to_string());
        params.insert("size",size.to_string());
        
        self.client.get(String::from("/pt/user/profile/invited"),params)
    }

	

	/**
	 * 
	 * 我的会员
	 */
    async fn myVip(&self, ) -> Result<ApiResultUserVipRechargeAccount,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/user/profile/vip"),params)
    }

	

	/**
	 * 
	 * 查询用户信息
	 */
    async fn profile(&self, ) -> Result<ApiResultUserPrincipalEntity,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/pt/user/profile"),params)
    }

	

	/**
	 * 
	 * 更新用户信息
	 */
    async fn update(&self, body: UserInitProfile) -> Result<ApiResultBoolean,Error>  {
        
        self.client.put(String::from("/pt/user/profile/update"),body)
    }

	
}











